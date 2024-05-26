package router

import (
	"backend/internal/handler/authhandler"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group) {
	g.POST("/login", authhandler.Login)
	g.POST("/signup", authhandler.Signup)
}
