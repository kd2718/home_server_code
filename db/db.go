package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Get_conn() *pgxpool.Pool {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		log.Fatalf("DB connection error: %v", err)
		os.Exit(1)
	}
	return pool
}
