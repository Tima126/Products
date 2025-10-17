package routes

import (
	"products/handlers"

	"github.com/go-chi/chi"
)

func InitProductRoutes(r chi.Router) {
	r.Get("/", handlers.ProductsPageHandler)
	r.Get("/products", handlers.ProductsAllHandler)
	r.Get("/products/{id}", handlers.ProductByIDHandler)
}
