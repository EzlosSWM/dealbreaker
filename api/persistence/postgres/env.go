package postgres

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	User     string
	Dbname   string
	Password string
}

func LoadEnv() (*DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &DBConfig{
		Host:     os.Getenv("PG_HOST"),
		User:     os.Getenv("PG_USER"),
		Dbname:   os.Getenv("PG_NAME"),
		Password: os.Getenv("PG_PASSWORD"),
	}, nil
}
