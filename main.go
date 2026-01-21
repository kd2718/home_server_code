package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

//var conn *pgx.Conn

const PORT = "8081"
const MAXBODY = 1024 * 1024

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	secretKey := os.Getenv("WEB_API_SECRET_KEY")
	if secretKey == "" {
		log.Fatalf("WEB_API_SECRET_KEY is not found. Must be set")
	}
	if len(secretKey) < 15 {
		log.Fatalf("WEB_API_SECRET_KEY is too short. Must be 15 or more characters long")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking Auth")
		apiKey := r.Header.Get("X-API-KEY")
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(secretKey)) != 1 {
			log.Printf("Request with bad API key")
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func getSensorData(dbpool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Post only", http.StatusMethodNotAllowed)
			return
		}

		// limit body size so that we don't read 10 GB of garbage data
		r.Body = http.MaxBytesReader(w, r.Body, MAXBODY)
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)

		if err != nil {
			// TODO fix
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		if !json.Valid(body) {
			log.Printf("Got bad json in body: %v\n", string(body))
			http.Error(w, "Bad JSON in body", http.StatusBadRequest)
			return
		}

		log.Printf("New Post: %v\n", string(body))
		// basic 201 return. doesn't care about body

		query := "INSERT INTO sensor_data (payload) values ($1)"
		_, err = dbpool.Exec(r.Context(), query, string(body))
		if err != nil {
			log.Printf("Insert failed: %v\n", err)
			http.Error(w, "Error connection to DB", http.StatusInternalServerError)
			return
		}
		log.Println("Row Inserted")
		w.WriteHeader(http.StatusCreated)
	}
}

func get_conn() *pgxpool.Pool {
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

func main() {
	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = PORT
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Port number is bad. Must be integer. got: %v\n", port)
	}
	serverPort := fmt.Sprintf(":%v", portInt)
	log.Printf("Now serving on: \"%v\"\n", serverPort)

	pool := get_conn()
	defer pool.Close()

	http.HandleFunc("/sensor", checkAuth(getSensorData(pool)))

	log.Fatal(http.ListenAndServe(serverPort, nil))
}
