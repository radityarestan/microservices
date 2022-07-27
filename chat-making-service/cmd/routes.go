package main

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Heartbeat("/ping"))
	mux.HandleFunc("/find-someone", app.findSomeone)

	return mux
}
