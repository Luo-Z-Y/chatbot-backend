package authhandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/auth"
	"backend/internal/database"
	authparams "backend/internal/params/authparams"
	"backend/internal/util"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	r := new(authparams.Params)
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db := database.GetDb()
	user, err := auth.CheckLogin(db, r.Username, r.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	t, err := util.GetJwtToken(user.ID)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	jwtCookie := http.Cookie{Name: "token", Value: t, Expires: expiration}

	c.SetCookie(&jwtCookie)
	return c.JSON(http.StatusOK, api.Response{Data: t})
}

func GetUser(c echo.Context) error {
	c.Logger().Printf("%+v", c.Get("user"))
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errors.New("JWT token missing or invalid"))
	}
	userModel, err := util.GetUser(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	userView := userModel.ToView()
	return c.JSON(http.StatusOK, api.Response{Data: userView})
}

func Signup(c echo.Context) error {
	r := new(authparams.Params)
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := r.ToModel()
	if err := user.Create(database.GetDb()); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
