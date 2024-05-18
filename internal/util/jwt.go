package util

import (
	"backend/internal/configs"
	"backend/internal/dataaccess/auth"
	"backend/internal/database"
	"backend/internal/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GetJwtToken(ID uint) (string, error) {
	claims := &JwtCustomClaims{ID: strconv.Itoa(int(ID)), RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.GetJwtSecret()))
}

func GetUser(jwt *jwt.Token) (*model.User, error) {
	claims := jwt.Claims.(*JwtCustomClaims)
	userId := claims.ID
	userIdInt, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return nil, err
	}

	userModel, err := auth.Read(database.GetDb(), uint(userIdInt))
	if err != nil {
		return nil, err
	}
	return userModel, nil
}
