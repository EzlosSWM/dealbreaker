package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"redCards/router"
	"redCards/storage"
)

func main() {

	if err := storage.NewPostgresStore(); nil != err {
		panic(err)
	}

	if err := storage.InitTable(); nil != err {
		panic(err)
	}

	app := router.StartApp()

	// app.Logger.Fatal(app.Start(":3000"))
	go func() {
		if err := app.Start(":3000"); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
