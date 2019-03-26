package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var port int

type key string

func main() {
	port = 3000
	fmt.Println("Connecting to port", port)
	r := chi.NewRouter()

	// r.Use(middleware.Logger)

	r.Get("/", landingPageHandler)

	r.Route("/scrape", func(r chi.Router) {
		r.Get("/*", reqHandler)
		r.Get("/", emptyHandler)
	})
	http.ListenAndServe(":3000", r)
}

// HTTP handler accessing the url routing parameters.
func reqHandler(w http.ResponseWriter, r *http.Request) {
	urlString := chi.URLParam(r, "*") // from a route like /scrape/{url}

	var urlKey key
	urlKey = "url"

	ctx := context.WithValue(r.Context(), urlKey, urlString)
	fmt.Println(ctx)

	w.Write([]byte(fmt.Sprintf("Scraping %v", ctx.Value(urlKey))))
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Please supply url to scrape"))
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to idgarza.dev kh scraper."))
}
