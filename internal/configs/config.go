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

	JwtSecret string `envconfig:"JWT_SECRET" default:"secret"`
}

var globalCfg *Config

func GetConfig() (*Config, error) {
	if globalCfg != nil {
		return globalCfg, nil
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := new(Config)
	err = envconfig.Process("", cfg)

	globalCfg = cfg
	return cfg, err
}

func GetJwtSecret() string {
	if globalCfg != nil {
		return globalCfg.JwtSecret
	} else {
		cfg, err := GetConfig()
		if err != nil {
			panic(err)
		}
		return cfg.JwtSecret
	}
}
