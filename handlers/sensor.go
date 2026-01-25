package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

const MAXBODY = 1024 * 1024

func GetSensorData(dbpool *pgxpool.Pool) http.HandlerFunc {
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
