package main

import (
	"backend/internal/configs"
	"backend/internal/database"
	"backend/internal/router"
	"backend/internal/telegram"
	"backend/internal/ws"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	database.SetupDb(cfg.GetDatabaseConfig())

	hub := ws.NewHub()
	go hub.Run()

	e := router.Setup(cfg, hub)
	go telegram.StartChatbot(cfg, hub)

	// Writing routes for debugging - we can optionally delete later
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	_ = os.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
