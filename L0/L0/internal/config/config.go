package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PgHost          string
	PgPort          string
	PgUser		    string
	PgPassword      string	
	PgDatabase      string

	NatsPort       string
	NatsCluster    string
	NatsPublisher  string
	NatsSubscriber string 
	NatsSubject    string

	ServerPort string
}

func InitConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can't loading .env")
	}

	cfg := &Config{
		PgHost: os.Getenv("PG_HOST"),
		PgPort: os.Getenv("PG_PORT"),		
		PgUser: os.Getenv("PG_USER"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgDatabase: os.Getenv("PG_DATABASE"),		
	
		NatsPort: os.Getenv("NATS_PORT"),
		NatsCluster: os.Getenv("NATS_CLUSTER"),
		NatsPublisher: os.Getenv("NATS_PUB"),
		NatsSubscriber: os.Getenv("NATS_SUB"),
		NatsSubject: os.Getenv("NATS_SUBJECT"),
	
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	return cfg
}

