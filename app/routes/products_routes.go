package routes

import (
	"products/handlers"

	"github.com/go-chi/chi"
)

func InitProductRoutes(r chi.Router) {
	r.Get("/", handlers.ProductsPageHandler)
	r.Get("/products", handlers.ProductsAllHandler)
	r.Get("/products/{id}", handlers.ProductByIDHandler)
	r.Post("/products", handlers.CreateNewProducts)
	r.Put("/products/{id}", handlers.UpdateProducts)
	r.Delete("/products/{id}", handlers.ProductDeleteOfid)

}
