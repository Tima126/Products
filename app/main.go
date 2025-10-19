package main

import (
	"fmt"
	"log"
	"net/http"
	"products/db"
	"products/routes"
)

func main() {

	db.InitPgxPool()
	r := routes.RegisterRoutes()
	defer db.Pool.Close()

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
