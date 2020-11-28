package server

import (
	"github.com/go-chi/chi"
	"go-hansolo/pkg/server/services"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/", services.GetLocation(500, 200, 1538))

	/*r.Post("/", ur.CreateHandler)

	r.Get("/{id}", ur.GetOneHandler)

	r.Put("/{id}", ur.UpdateHandler)

	r.Delete("/{id}", ur.DeleteHandler)*/

	return r
}