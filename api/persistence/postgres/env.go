package postgres

import (
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	Dbname   string
	Password string
}

func LoadEnv() (*DBConfig, error) {
	return &DBConfig{
		Host:     os.Getenv("HOSTADDR"),
		User:     os.Getenv("USER_NAME"),
		Dbname:   os.Getenv("DB_NAME"),
		Password: os.Getenv("PASSWORD"),
	}, nil
}
