package database

import (
	"backend/internal/configs"
	"fmt"
	"strings"
)

func BuildDsn(cfg *configs.PostgresConfig) (string, error) {
	dnsBuilder := strings.Builder{}

	_, err := dnsBuilder.WriteString(fmt.Sprintf("host=%s", cfg.PostgresHost))
	if err != nil {
		return "", err
	}

	_, err = dnsBuilder.WriteString(fmt.Sprintf(" user=%s", cfg.PostgresUser))
	if err != nil {
		return "", err
	}

	_, err = dnsBuilder.WriteString(fmt.Sprintf(" password=%s", cfg.PostgresPassword))
	if err != nil {
		return "", err
	}

	_, err = dnsBuilder.WriteString(fmt.Sprintf(" dbname=%s", cfg.PostgresDb))
	if err != nil {
		return "", err
	}

	_, err = dnsBuilder.WriteString(fmt.Sprintf(" port=%s", cfg.PostgresPort))
	if err != nil {
		return "", err
	}

	_, err = dnsBuilder.WriteString(" sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		return "", err
	}

	return dnsBuilder.String(), nil
}
