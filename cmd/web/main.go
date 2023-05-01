package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/M0F3/bedandbreakfast/pkg/config"
	"github.com/M0F3/bedandbreakfast/pkg/handlers"
	"github.com/M0F3/bedandbreakfast/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8082"
var app config.AppConfig
var session * scs.SessionManager

// main is the Main function
func main() {


	//Has to be true in production environment
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteStrictMode

	session.Cookie.Secure = app.InProduction

	app.Session = session

	ts, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = ts
	app.UseCache = true

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	fmt.Printf("Start application on port %s\n", portNumber)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}