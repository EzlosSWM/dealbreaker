package handlers

import (
	"fmt"
	"net/http"
	"redCards/internal/app/domain"
	"redCards/internal/app/service"
	"redCards/internal/infastructure/models"
	"redCards/internal/infastructure/persistence/postgres"

	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

// create card
func CreateCard(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	result, err := domain.Create(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
	}

	return e.JSON(http.StatusCreated, result)
}

// get all jokes and/or filters
func GetJokes(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	jokeType := e.QueryParam("joke_type")
	result, err := domain.Get(jokeType)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, result)
}

// delete
func DeleteJoke(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
	}

	if err = domain.Delete(id); err != nil {
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, fmt.Sprintf("Card deleted: %d", id))
}

func BatchUpload(e echo.Context) error {
	var cards []models.Card
	err := e.Bind(&cards)
	if err != nil {
		return e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
	}

	// repository
	for _, card := range cards {
		_, err := postgres.DB.Exec(`INSERT INTO cards (joke_type, joke) VALUES ($1, $2)`, card.JokeType, card.Joke)
		if err != nil {
			return err
		}
	}

	return e.JSON(http.StatusCreated, cards)
}
