package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/kennnyz/bookings/pckg/config"
	Handler "github.com/kennnyz/bookings/pckg/handler"
	"github.com/kennnyz/bookings/pckg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(" Can not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := Handler.NewRepo(&app)
	Handler.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", Handler.Repo.Home)
	//http.HandleFunc("/about", Handler.Repo.About)

	fmt.Printf("Staring application on port %s", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
