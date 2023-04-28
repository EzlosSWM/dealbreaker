package controller

import (
	"database/sql"
	"net/http"
	"redCards/models"
	"redCards/storage"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

// create perk
func CreatePerk(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	perk := new(models.Card)
	if err := e.Bind(&perk); err != nil {
		return err
	}

	perk.JokeType = "perk"

	res, err := storage.DB.Query("INSERT INTO cards (joke_type, joke, topic)VALUES ($1,$2,$3)", perk.JokeType, perk.Joke, perk.Topic)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error: ": err.Error()})
	}

	return e.JSON(http.StatusCreated, res)
}

// create new deal breaker
func CreateDealbreaker(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	dealbreaker := new(models.Card)
	if err := e.Bind(&dealbreaker); err != nil {
		return err
	}

	dealbreaker.JokeType = "dealbreaker"

	res, err := storage.DB.Query("INSERT INTO cards (joke_type, joke, topic)VALUES ($1,$2,$3)", dealbreaker.JokeType, dealbreaker.Joke, dealbreaker.Topic)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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
		return err
	}

	result, err := scanIntoCard(rows)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, result)
}

// delete
func DeleteJoke(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	_, err = storage.DB.Query("DELETE FROM cards WHERE id = $1", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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
