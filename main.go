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

	r.Route("/login", func(r chi.Router) {
		r.Get("/{user}", reqHandler)
		r.Get("/", emptyHandler)
	})
	http.ListenAndServe(":3000", r)
}

// HTTP handler accessing the url routing parameters.
func reqHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user") // from a route like /login/{user}

	var userKey key
	userKey = "userName"

	ctx := context.WithValue(r.Context(), userKey, userID)

	w.Write([]byte(fmt.Sprintf("Logged in as %v", ctx.Value(userKey))))
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Please supply username"))
}
