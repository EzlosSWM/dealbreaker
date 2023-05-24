package routes

import (
	"net/http"
	"os"
	"redCards/api/handlers"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	jwt.RegisteredClaims
}

func Routes(e *echo.Echo) {
	secret := os.Getenv("JWT")

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	// grouping
	group := e.Group("/api/v1")
	group.GET("/jokes", handlers.GetJokes)

	// restricted routes
	r := group.Group("/card")

	// custom claims
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(secret),
	}
	r.Use(echojwt.WithConfig(config))

	r.POST("", handlers.CreateCard)
	r.POST("/batch", handlers.BatchUpload)
	r.DELETE("/:id", handlers.DeleteJoke)
}
