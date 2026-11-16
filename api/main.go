package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	r.Mount("/echoes", EchoRoutes())
	
	log.Fatal(http.ListenAndServe(":3000", r))
}


func EchoRoutes() chi.Router {
	eRouter := chi.NewRouter()

	echoHandler := EchoHandler{}
	eRouter.Get("/", echoHandler.ListEchoes)

	return eRouter
}
