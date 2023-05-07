package handlers

import (
	"net/http"
	"redCards/internal/app/domain"
	"redCards/internal/app/service"

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

// func TestCards(e echo.Context) error {
// 	var cards []models.Card
// 	err := e.Bind(&cards)
// 	if err != nil {
// 		return e.JSON(http.StatusBadRequest, service.ErrToJSON(err))
// 	}

// 	return e.JSON(http.StatusCreated, cards)
// }

// get all jokes and/or filters
func GetJokes(e echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	result, err := domain.Get(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
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
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	if err := domain.Delete(e, id); err != nil {
		return e.JSON(http.StatusInternalServerError, service.ErrToJSON(err))
	}

	return e.JSON(http.StatusOK, map[string]int{"deleted card with id: ": id})
}
