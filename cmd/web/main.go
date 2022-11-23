package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/handlers"
	"github.com/vikashparashar/bookings/pkg/render"
)

const (
	portNumber string = ":8080"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	// adding session's for server
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // true in production mode else false in developement mode
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("can not create template")
	}

	app.TemplateCache = tc
	app.UseCache = false // in development mode useCache need to be False otherwise in production mode it sets to be True
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
