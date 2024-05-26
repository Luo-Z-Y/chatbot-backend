package chathandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	db := database.GetDb()
	chats, err := chat.List(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to get chats", err.Error())
	}

	return c.JSON(http.StatusOK, chats)
}

func Read(c echo.Context) error {
	chatIDStr := c.Param("chat_id")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid chat_id "+err.Error())
	}

	db := database.GetDb()
	chat, err := chat.Read(db, uint(chatID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to get chat with id %d, %s", chatID, err.Error())
	}


	return c.JSON(http.StatusOK, api.Response{Data: })
	return c.JSON(http.StatusOK, chats)
}
