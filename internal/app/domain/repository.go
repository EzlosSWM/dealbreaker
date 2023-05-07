package domain

import (
	"database/sql"
	"net/http"
	"redCards/internal/app/service"
	"redCards/internal/infastructure/models"
	"redCards/internal/infastructure/persistence/postgres"

	"github.com/labstack/echo/v4"
)

func Get(e echo.Context) (*models.Cards, error) {
	var query string
	var args []interface{}

	if jokeType := e.QueryParam("joke_type"); jokeType != "" {
		query = "SELECT * FROM cards WHERE joke_type = $1"
		args = append(args, jokeType)
	} else {
		query = "SELECT * FROM cards"
	}

	if topic := e.QueryParam("topic"); topic != "" {
		if len(args) == 0 {
			query = "SELECT * FROM cards WHERE topic = $1"
		} else {
			query += " AND topic = $2"
		}

		args = append(args, topic)
	}

	rows, err := postgres.DB.Query(query, args...)
	if err != nil {
		return nil, e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	result, err := scanIntoCard(rows)
	if err != nil {
		return nil, e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	return result, nil
}

func Create(e echo.Context) (*sql.Rows, error) {
	card := new(models.Card)
	if err := e.Bind(&card); err != nil {
		return nil, err
	}

	result, err := postgres.DB.Query("INSERT INTO cards (joke_type, joke, topic)VALUES ($1,$2,$3)", card.JokeType, card.Joke, card.Topic)
	if err != nil {
		return nil, e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
	}

	return result, nil
}

func Delete(e echo.Context, id int) error {
	_, err := postgres.DB.Query("DELETE FROM cards WHERE id = $1", id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	return err
}

// scan functionality for getting cards from db
func scanIntoCard(rows *sql.Rows) (*models.Cards, error) {
	result := new(models.Cards)

	for rows.Next() {
		card := new(models.Card)
		if err := rows.Scan(&card.ID, &card.JokeType, &card.Joke, &card.Topic); err != nil {
			return nil, err
		}

		result.Cards = append(result.Cards, *card)
	}

	return result, nil
}
