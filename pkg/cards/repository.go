package cards

import (
	"database/sql"
	"log"
	"net/http"
	"redCards/api/persistence/postgres"
	"redCards/pkg/entities"

	"github.com/labstack/echo/v4"
)

func Get(jokeType string) (*entities.Cards, error) {
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
	card := new(entities.Card)
	if err := e.Bind(&card); err != nil {
		return nil, err
	}

	result, err := postgres.DB.Query("INSERT INTO cards (joke_type, joke) VALUES ($1,$2)", card.JokeType, card.Joke)
	if err != nil {
		return nil, e.JSON(http.StatusBadRequest, ErrToJSON(err))
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
func scanIntoCard(rows *sql.Rows) (*entities.Cards, error) {
	result := new(entities.Cards)

	for rows.Next() {
		card := new(entities.Card)
		if err := rows.Scan(&card.ID, &card.JokeType, &card.Joke); err != nil {
			return nil, err
		}

		result.Cards = append(result.Cards, *card)
	}

	return result, nil
}
