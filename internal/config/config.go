package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_CONNECTION  string
	JWT_SECRET_KEY string
	SENDER         string
}

var Env = initEnv()

func initEnv() Config {
	godotenv.Load()

	return Config{DB_CONNECTION: getEnv("DB_CONNECTION"), JWT_SECRET_KEY: getEnv("JWT_SECRET_KEY"), SENDER: getEnv("SENDER")}
}

func getEnv(key string) string {

	v, ok := os.LookupEnv(key)

	if !ok {
		return ""
	}
	return v
}
