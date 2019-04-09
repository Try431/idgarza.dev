package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Try431/idgarza.dev/scrape"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var port int

type key string

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/scrape", scrape.Routes())
	})
	return r
}

func main() {
	port = 3000
	fmt.Println("Connecting to port", port)
	router := Routes()
	log.Fatal(http.ListenAndServe(":3000", router))
}
