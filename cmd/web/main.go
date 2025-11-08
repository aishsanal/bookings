package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aishsanal/bookings/pkg/handlers"
	"github.com/aishsanal/bookings/pkg/config"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var (
	app     config.AppConfig
	session *scs.SessionManager
)

func main() {
	app.IsProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.IsProd
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	ts, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot start application")
	}

	app.TemplateCache = ts
	app.UseCache = false
	handlers.SetConfig(app)

	repository := *handlers.CreateRepository(app)
	handlers.SetRepository(&repository)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
