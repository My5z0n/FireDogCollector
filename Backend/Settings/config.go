package Settings

import "os"

type Config struct {
	EnvType     string
	BackendUrl  string
	BackendPort string
	DbUrl       string
	DbPort      string
}

func GetEnvConfig() Config {
	var c Config

	c.BackendUrl = os.Getenv("BACKEND_URL")
	c.BackendPort = os.Getenv("BACKEND_PORT")
	c.DbUrl = os.Getenv("DB_URL")
	c.DbPort = os.Getenv("DB_PORT")
	c.EnvType = os.Getenv("ENV_TYPE")

	return c
}
