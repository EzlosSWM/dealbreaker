package controller

import (
	"database/sql"
	"net/http"
	"redCards/helpers"
	"redCards/models"
	"redCards/storage"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

// create card
func CreateCard(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	card := new(models.Card)
	if err := e.Bind(&card); err != nil {
		return err
	}

	res, err := storage.DB.Query("INSERT INTO cards (joke_type, joke, topic)VALUES ($1,$2,$3)", card.JokeType, card.Joke, card.Topic)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error: ": err.Error()})
	}

	return e.JSON(http.StatusCreated, res)
}

// get all jokes and/or filters
func GetJokes(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

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

	rows, err := storage.DB.Query(query, args...)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helpers.ErrToJSON(err))
	}

	result, err := scanIntoCard(rows)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helpers.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, result)
}

// delete
// TODO: Check if ID exsists
func DeleteJoke(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helpers.ErrToJSON(err))
	}

	_, err = storage.DB.Query("DELETE FROM cards WHERE id = $1", id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helpers.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, map[string]int{"deleted card with id: ": id})
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
