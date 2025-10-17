package handlers

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"products/db"
	"strconv"

	"github.com/go-chi/chi"
)

// структура продукта
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ProductsPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/products_all.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// вывод всех продуктов
func ProductsAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rows, err := db.Pool.Query(ctx, "SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		log.Println("DB error:", err)
		return
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			http.Error(w, "Failed to scan product", http.StatusInternalServerError)
			log.Println("DB error:", err)
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при чтении результатов", http.StatusInternalServerError)
		log.Println("Rows error:", err)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func ProductByIDHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var p Product

	err := db.Pool.QueryRow(ctx,
		"SELECT id, name, description, price FROM products WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(p)

}
