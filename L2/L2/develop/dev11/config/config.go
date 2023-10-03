package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}
//Получаем порт и хост из .env
func InitConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can't loading .env: ", err)
	}

	cfg := &Config{
		Host:     os.Getenv("Host"),
		Port:     os.Getenv("Port"),
	}

	return cfg
}
