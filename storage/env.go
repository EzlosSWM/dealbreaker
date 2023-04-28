package storage

import (
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
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &DBConfig{
		Host:     os.Getenv("HOSTADDR"),
		User:     os.Getenv("USER_NAME"),
		Dbname:   os.Getenv("DB_NAME"),
		Password: os.Getenv("PASSWORD"),
	}, nil
}
