package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_CONNECTION string
}

var Env = initEnv()

func initEnv() Config {
	godotenv.Load()

	return Config{DB_CONNECTION: getEnv("DB_CONNECTION")}
}

func getEnv(key string) string {

	v, ok := os.LookupEnv(key)

	if !ok {
		return ""
	}
	return v
}
