package middleware

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
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
