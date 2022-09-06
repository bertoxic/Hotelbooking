package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
	"github.com/bertoxic/bookings/pkg/config"
	"github.com/bertoxic/bookings/pkg/handlers"
	"github.com/bertoxic/bookings/pkg/render"
)

var portNumber = ":8080"
var Session *scs.SessionManager
var app config.AppConfig



func main() {
		
	

	app.InProduction=false

	Session=scs.New()
	Session.Lifetime = 26* time.Hour
	Session.Cookie.Persist = true
	Session.Cookie.SameSite = http.SameSiteLaxMode
	Session.Cookie.Secure= app.InProduction
	app.Session=Session	
	

	tc, err := render.CreateTemplate()
	if err != nil {
		log.Fatal("error in main maybe", err)
	}
	fmt.Println("check if we empty here...",tc)
	app.TemplateCache = tc	
	fmt.Println(len(tc))
	app.UserCache=false
	

	repo:=handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
	Addr:portNumber,
	Handler: routes(&app),

	}
	
	err =srv.ListenAndServe()
	log.Fatal("okkkkkkkkkk in main",err)


}
