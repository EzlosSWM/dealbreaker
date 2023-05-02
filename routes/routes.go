package routes

import (
	"net/http"
	"os"
	"redCards/controller"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	jwt.RegisteredClaims
}

func StartApp() *echo.Echo {
	secret := os.Getenv("JWT")

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
	group.GET("/jokes", controller.GetJokes)

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

	r.POST("", controller.CreateCard)
	r.DELETE("/:id", controller.DeleteJoke)

	return e
}
