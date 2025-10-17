package main

import (
	"fmt"
	"net/http"
	"products/db"
	"products/routes"
)

func main() {

	db.InitPgxPool()
	r := routes.RegisterRoutes()
	defer db.Pool.Close()

	fmt.Println(http.ListenAndServe("localhost:8181", r))

}
