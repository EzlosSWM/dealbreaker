package domain

import (
	"database/sql"
	"log"
	"net/http"
	"redCards/internal/app/service"
	"redCards/internal/infastructure/models"
	"redCards/internal/infastructure/persistence/postgres"

	"github.com/labstack/echo/v4"
)

func Get(jokeType string) (*models.Cards, error) {
	query := "SELECT * FROM cards"

	// jokeType := e.QueryParam("joke_type")
	if jokeType != "" {
		query += " WHERE joke_type = $1"
	}

	var rows *sql.Rows
	var err error
	if jokeType != "" {
		rows, err = postgres.DB.Query(query, jokeType)
	} else {
		rows, err = postgres.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	log.Printf("%v", query)

	result, err := scanIntoCard(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Create(e echo.Context) (*sql.Rows, error) {
	card := new(models.Card)
	if err := e.Bind(&card); err != nil {
		return nil, err
	}

	result, err := postgres.DB.Query("INSERT INTO cards (joke_type, joke) VALUES ($1,$2)", card.JokeType, card.Joke)
	if err != nil {
		return nil, e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
	}

	return result, nil
}

func BatchPost(e echo.Context) error {

	return nil
}

func Delete(id int) error {
	res, err := postgres.DB.Exec("DELETE FROM cards WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}

// scan functionality for getting cards from db
func scanIntoCard(rows *sql.Rows) (*models.Cards, error) {
	result := new(models.Cards)

	for rows.Next() {
		card := new(models.Card)
		if err := rows.Scan(&card.ID, &card.JokeType, &card.Joke); err != nil {
			return nil, err
		}

		result.Cards = append(result.Cards, *card)
	}

	return result, nil
}
