package main

import (
	"net/http"

	"github.com/aishsanal/bookings/pkg/config"
	"github.com/aishsanal/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(printToConsole)
	mux.Use(NoSurf)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)
	mux.Get("/thumpa", handlers.Thumpa)
	mux.Get("/mulla", handlers.Mulla)
	mux.Get("/make-reservation", handlers.Reservation)

	mux.Get("/check-availability", handlers.Availability)
	mux.Post("/check-availability", handlers.PostAvailability)
	mux.Get("/check-availability-json", handlers.JSONPostAvailability)

	mux.Get("/contact", handlers.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
