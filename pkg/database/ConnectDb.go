package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDb() *pgxpool.Pool {
	connString := os.Getenv("POSTGRES_CONNECTION_STRING")
	pool, err := pgxpool.Connect(context.TODO(), connString)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}