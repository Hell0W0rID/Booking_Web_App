package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Hell0W0rID/Booking_Web_App/pkg/handlers"
	"github.com/Hell0W0rID/Booking_Web_App/pkg/handlers/config"
	render "github.com/Hell0W0rID/Booking_Web_App/pkg/handlers/renders"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// })

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true            //
	session.Cookie.Secure = app.InProduction // for encrypted connection

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.ListenAndServe(":8081", nil)

}
