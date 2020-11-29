package controller

import (
	"github.com/go-chi/chi"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	service := &Impl{}
	r.Mount("/hansolo", service.Routes())

	return r
}