package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GoEnv string `envconfig:"GO_ENV" default:"development"`

	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser     string `envconfig:"POSTGRES_USER"  default:"postgres"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresDb       string `envconfig:"POSTGRES_DB" default:"chatbot"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`

	TestPostgresHost     string `envconfig:"TEST_POSTGRES_HOST" default:"localhost"`
	TestPostgresUser     string `envconfig:"TEST_POSTGRES_USER"  default:"postgres"`
	TestPostgresPassword string `envconfig:"TEST_POSTGRES_PASSWORD" default:"postgres"`
	TestPostgresDb       string `envconfig:"TEST_POSTGRES_DB" default:"chatbot"`
	TestPostgresPort     string `envconfig:"TEST_POSTGRES_PORT" default:"5434"`

	Port          string `envconfig:"PORT" default:"8000"`
	TelegramToken string `envconfig:"TELEGRAM_TOKEN" required:"true"`
	FrontendUrl   string `envconfig:"FRONTEND_URL" default:"http://localhost:3000"`

	JwtSecret string `envconfig:"JWT_SECRET" default:"secret"`
}

type PostgresConfig struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresPort     string
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

func (c *Config) GetDatabaseConfig() *PostgresConfig {
	return &PostgresConfig{
		c.PostgresHost,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDb,
		c.PostgresPort,
	}
}

func (c *Config) GetTestDatabaseConfig() *PostgresConfig {
	return &PostgresConfig{
		c.TestPostgresHost,
		c.TestPostgresUser,
		c.TestPostgresPassword,
		c.TestPostgresDb,
		c.TestPostgresPort,
	}
}
