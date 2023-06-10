package main

import (
	"net/http"
	"redCards/api/persistence/postgres"
	"redCards/api/routes"
)

func main() {

	if err := postgres.NewPostgresStore(); nil != err {
		panic(err)
	}

	if err := postgres.InitTable(); nil != err {
		panic(err)
	}

	app := routes.Start()

	// if err := routes.Shutdown(app); nil != err {
	// 	panic(err)
	// }

	var port string
	port = "3000"

	if err := app.Start(":" + port); err != nil && err != http.ErrServerClosed {
		app.Logger.Fatal("shutting down the server")
	}
}
