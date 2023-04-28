package router

import (
	"net/http"
	"redCards/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartApp() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	// grouping
	group := e.Group("/api/v1")

	// routes
	// create
	group.POST("/perk", controller.CreatePerk)
	group.POST("/dealbreaker", controller.CreateDealbreaker)

	// read
	group.GET("/jokes", controller.GetJokes)

	// delete
	group.DELETE("/jokes/:id", controller.DeleteJoke)

	return e
}
