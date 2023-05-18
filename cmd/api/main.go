package main

import (
	"redCards/internal/infastructure/http"
	"redCards/internal/infastructure/persistence/postgres"
)

func main() {

	if err := postgres.NewPostgresStore(); nil != err {
		panic(err)
	}

	if err := postgres.InitTable(); nil != err {
		panic(err)
	}

	app := http.Start()

	if err := http.Shutdown(app); nil != err {
		panic(err)
	}
}
