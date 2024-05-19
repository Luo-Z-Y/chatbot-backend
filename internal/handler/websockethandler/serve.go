package websockethandler

import (
	"backend/internal/ws"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub *ws.Hub) func(c echo.Context) error {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		hub.Register <- conn

		go func() {
			defer func() {
				hub.Unregister <- conn
			}()
			for {
				// At the moment, we are only emitting messages to the client and not receiving any
				_, _, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("unexpected close error: %v", err)
					}
					break
				}
			}
		}()

		return nil
	}
}
