package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafixcs/uberemailservicechallenge/src/adapters/factories"
)

func LoadEmailRoutes(r chi.Router) {
	r.Post("/email", func(w http.ResponseWriter, r *http.Request) {
		controller := factories.NewEmailController()
		controller.Handle(w, r)
	})
}
