package routes

import (
	"github.com/go-chi/chi"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	InitProductRoutes(r)
	return r
}
