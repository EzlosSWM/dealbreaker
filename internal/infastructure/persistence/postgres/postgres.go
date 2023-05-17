package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewPostgresStore() error {
	var err error
	dbconfig, err := LoadEnv()
	if err != nil {
		return err
	}

	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbconfig.Host, dbconfig.User, dbconfig.Dbname, dbconfig.Password)
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err

	}

	return nil
}

func InitTable() error {
	query := `CREATE TABLE IF NOT EXISTS cards (
		id serial primary key,
		joke_type text not null,
		joke text not null
	)
	`
	_, err := DB.Exec(query)

	return err
}
