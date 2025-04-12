package main

import (
	"inky_go/templates"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("./static"))

	r.Use(middleware.Logger)
	r.Get("/", templ.Handler(templates.Page(getTime())).ServeHTTP)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("listening on port :3000")
	http.ListenAndServe(":3000", r)
}

func getTime() string {
	time := time.Now()
	return time.Format("15:04:05")
}
