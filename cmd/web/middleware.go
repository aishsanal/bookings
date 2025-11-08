package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func printToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

//NoSurf provides CSRF protection for all requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure: false,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//Loads and saves the session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
