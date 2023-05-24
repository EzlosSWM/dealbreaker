package main

import (
	"net/http"
	"redCards/api/persistence/postgres"
	"redCards/api/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {

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

	app.ServeHTTP(w, r)
}
