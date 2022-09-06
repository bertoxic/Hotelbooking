package main

import (
	"net/http"

	"github.com/bertoxic/bookings/pkg/config"
	"github.com/bertoxic/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/Majors", handlers.Repo.Majors)
	mux.Get("/Reservations", handlers.Repo.Reservations)
	mux.Post("/Reservations", handlers.Repo.PostReservations)
	mux.Get("/makeReservation", handlers.Repo.MakeReservation)
	mux.Get("/AvailabilityJson", handlers.Repo.AvailabilityJson)

	fileServer := http.FileServer(http.Dir("../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
