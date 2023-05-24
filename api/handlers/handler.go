package handlers

import (
	"fmt"
	"net/http"
	"redCards/api/persistence/postgres"
	"redCards/pkg/cards"
	"redCards/pkg/entities"

	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

// create card
func CreateCard(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	result, err := cards.Create(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, cards.ErrToJSON(err))
	}

	return e.JSON(http.StatusCreated, result)
}

// get all jokes and/or filters
func GetJokes(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	jokeType := e.QueryParam("joke_type")
	result, err := cards.Get(jokeType)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, cards.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, result)
}

// delete
func DeleteJoke(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, cards.ErrToJSON(err))
	}

	if err = cards.Delete(id); err != nil {
		return e.JSON(http.StatusInternalServerError, cards.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, fmt.Sprintf("Card deleted: %d", id))
}

func BatchUpload(e echo.Context) error {
	var cardsUp []entities.Card
	err := e.Bind(&cardsUp)
	if err != nil {
		return e.JSON(http.StatusBadRequest, cards.ErrToJSON(err))
	}

	// repository
	for _, card := range cardsUp {
		_, err := postgres.DB.Exec(`INSERT INTO cards (joke_type, joke) VALUES ($1, $2)`, card.JokeType, card.Joke)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusCreated, cardsUp)
}
