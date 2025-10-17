package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPgxPool() {

	dsn := "postgres://admin:12345@localhost:5435/product_db?sslmode=disable"

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	log.Println("Connection pool created successfully")

}
