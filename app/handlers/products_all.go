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
	tmpl := template.Must(template.ParseFiles("templates/products_all.html"))
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

// вспомогательная функция для получения id
func GetIdProducts(ctx context.Context, id int) (*Product, error) {
	var p Product
	err := db.Pool.QueryRow(ctx,
		"SELECT id, name, description, price FROM products WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil

}

// получени продуктов по id
func ProductByIDHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	p, err := GetIdProducts(ctx, id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

// удаление записи
func ProductDeleteOfid(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	_, err = GetIdProducts(ctx, id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	_, err = db.Pool.Exec(ctx, "DELETE FROM products WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

}

func CreateNewProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price <= 0 {
		http.Error(w, "Name and positive price are required", http.StatusBadRequest)
		return
	}

	err := db.Pool.QueryRow(ctx,
		"INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id",
		p.Name, p.Description, p.Price,
	).Scan(&p.ID)

	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 create
	json.NewEncoder(w).Encode(p)

}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	idstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price <= 0 {
		http.Error(w, "Name and positive price are required", http.StatusBadRequest)
		return
	}

	result, err := db.Pool.Exec(ctx,
		"UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4",
		p.Name, p.Description, p.Price, id,
	)

	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected() == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	p.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
