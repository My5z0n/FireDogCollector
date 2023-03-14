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

	c.BackendUrl = SetDef(os.Getenv("BACKEND_URL"), "localhost")
	c.BackendPort = SetDef(os.Getenv("BACKEND_PORT"), "9900")
	c.DbUrl = SetDef(os.Getenv("DB_URL"), "localhost")
	c.DbPort = SetDef(os.Getenv("DB_PORT"), "9000")
	c.EnvType = SetDef(os.Getenv("ENV_TYPE"), "")

	return c
}

func SetDef(env string, def string) string {
	if env == "" {
		return def
	} else {
		return env
	}
}
