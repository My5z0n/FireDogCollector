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

	c.OtelUrl = SetOrDefault(os.Getenv("OTEL_COLLECTOR_URL"), "localhost")
	c.OtelPort = SetOrDefault(os.Getenv("OTEL_COLLECTOR_PORT"), "4320")
	c.DbUrl = SetOrDefault(os.Getenv("DB_URL"), "localhost")
	c.DbPort = SetOrDefault(os.Getenv("DB_PORT"), "9000")

	return c
}

func SetOrDefault(env string, def string) string {
	if env != "" {
		return env
	} else {
		return def
	}
}
