package util

import (
	"backend/internal/model"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	User *model.User
}
