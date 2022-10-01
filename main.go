package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hugovallada/go-api-postgres/configs"
	"github.com/hugovallada/go-api-postgres/handlers"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := routes()

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

func routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/todos", handlers.Create)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)
	r.Get("/todos", handlers.GetAll)
	r.Get("/todos/{id}", handlers.Get)
	return r
}
