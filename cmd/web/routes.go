package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vikashparashar/Hotel_Bookings_2/pkg/config"
	"github.com/vikashparashar/Hotel_Bookings_2/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	// Using Chi Routing and Middelware

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	// mux.Use(WriteToConsole)

	// Nosurf package is used for stoping CSRf (cross site ref )
	mux.Use(NoSurf)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/general", http.HandlerFunc(handlers.Repo.General))
	mux.Get("/major", http.HandlerFunc(handlers.Repo.Major))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/check", http.HandlerFunc(handlers.Repo.CheckAvailability))
	fileserver := http.FileServer(http.Dir("./htmlfiles/static/"))
	mux.Handle("/htmlfiles/static/*", http.StripPrefix("/htmlfiles/static", fileserver))
	return mux

}
