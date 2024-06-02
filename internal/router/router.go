package router

import (
	"backend/internal/configs"
	"backend/internal/handler/authhandler"
	"backend/internal/handler/bothandler"
	"backend/internal/handler/websockethandler"
	"backend/internal/util"
	"backend/internal/ws"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Setup(cfg *configs.Config, hub *ws.Hub) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.FrontendUrl},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	g := e.Group("/api")

	// Routes that don't need authentication
	AuthRoutes(g)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JwtCustomClaims)
		},
		SigningKey:  []byte(cfg.JwtSecret),
		TokenLookup: "header:Authorization:Bearer ,cookie:token",
	}
	g.Use(echojwt.WithConfig(config))

	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return c.JSON(http.StatusUnauthorized, errors.New("JWT token missing or invalid"))
			}
			userModel, err := util.GetUser(user)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			cc := &util.CustomContext{Context: c, User: userModel}
			return next(cc)
		}
	})

	// Routes needing, authentication
	chatGroup := g.Group("/chats")
	ChatRoutes(chatGroup)

	g.GET("/current-user", authhandler.GetUser)

	botg := g.Group("/bot")
	botg.POST("/send/:chat_id", bothandler.SendMessage)

	tg := g.Group("/tg")
	tg.POST("/auth", authhandler.AuthenticateTgUser)

	// Websocket
	g.GET("/ws", websockethandler.ServeWs(hub))

	return e
}
