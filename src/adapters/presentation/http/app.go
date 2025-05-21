package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafixcs/uberemailservicechallenge/src/adapters/presentation/routes"
)

const PORT = "3000"

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.LoadRoutes(r)

	log.Printf("Server running on PORT: %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
