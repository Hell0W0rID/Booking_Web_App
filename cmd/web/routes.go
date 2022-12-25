package main

import (
	"net/http"

	"github.com/Hell0W0rID/Booking_Web_App/pkg/handlers"
	"github.com/Hell0W0rID/Booking_Web_App/pkg/handlers/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
