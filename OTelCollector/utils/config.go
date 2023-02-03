package utils

import "os"

type Config struct {
	OtelUrl  string
	OtelPort string
	DbUrl    string
	DbPort   string
}

func GetEnvConfig() Config {
	var c Config

	c.OtelUrl = os.Getenv("OTEL_COLLECTOR_URL")
	c.OtelPort = os.Getenv("OTEL_COLLECTOR_PORT")
	c.DbUrl = os.Getenv("DB_URL")
	c.DbPort = os.Getenv("DB_PORT")

	return c
}
