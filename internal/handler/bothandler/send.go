package bothandler

import (
	"backend/internal/params/tgmessageparams"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SendMessage(c echo.Context) error {
	chatIDStr := c.Param("chat_id")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid chat_id "+err.Error())
	}
	r := new(tgmessageparams.MessageParams)
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = SendTelegramMessage(int64(chatID), nil, r.Message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Message sent")
}
