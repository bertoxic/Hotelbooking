package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func writetoconsole(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("print smthing to console")
		 next.ServeHTTP(w, r)
	})

	
}

func NoSurf(next http.Handler) http.Handler{

	csrfHandler:= nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler{
	return Session.LoadAndSave(next)
}