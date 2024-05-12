package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser     string `envconfig:"POSTGRES_USER"  default:"postgres"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresDb       string `envconfig:"POSTGRES_DB" default:"chatbot"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`

	Port          string `envconfig:"PORT" default:"8000"`
	TelegramToken string `envconfig:"TELEGRAM_TOKEN" required:"true"`
	FrontendUrl   string `envconfig:"FRONTEND_URL" default:"http://localhost:3000"`
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := new(Config)
	err = envconfig.Process("", cfg)
	return cfg, err
}
