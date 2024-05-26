package router

import (
	"backend/internal/handler/chathandler"

	"github.com/labstack/echo/v4"
)

func ChatRoutes(g *echo.Group) {
	g.GET("/:chat_id", chathandler.Read)
	g.GET("", chathandler.List)
}
