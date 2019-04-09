package scrape

import (
	"fmt"
	"net/http"

	"github.com/Try431/idgarza.dev/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type key string

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/*", scrapeURL)
	r.Get("/", emptyHandler)
	return r
}

// HTTP handler accessing the url routing parameters.
func scrapeURL(w http.ResponseWriter, r *http.Request) {
	urlString := chi.URLParam(r, "*") // from a route like /scrape/{url}
	fmt.Println(urlString)
	// var urlKey key
	// urlKey = "url"

	// ctx := context.WithValue(r.Context(), urlKey, urlString)
	// fmt.Println(ctx)
	render.JSON(w, r, "hey")

	html := grabHTML(urlString)
	resp := model.ResponseStruct{
		RequestURL:   urlString,
		ResponseHTML: html}

	render.JSON(w, r, resp)

	// w.Write([]byte(fmt.Sprintf("Scraping %v", ctx.Value(urlKey))))
	w.Write([]byte(fmt.Sprintf("Scraping %v", urlString)))
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Please supply url to scrape"))
}

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to idgarza.dev kh scraper."))
}
