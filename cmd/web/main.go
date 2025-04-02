package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Prateek766/bookings/internal/config"
	"github.com/Prateek766/bookings/internal/handlers"
	"github.com/Prateek766/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":4000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Application running on port %s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
