package routes

import (
	"github.com/go-chi/chi/v5"
	routes "github.com/rafixcs/uberemailservicechallenge/src/adapters/presentation/routes/email"
)

func LoadRoutes(r *chi.Mux) {
	routes.LoadEmailRoutes(r)
}
