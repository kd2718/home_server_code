package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"server/hsc/db"
	"server/hsc/handlers"
	"server/hsc/middleware"
)

//var conn *pgx.Conn

const PORT = "8081"

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

	pool := db.Get_conn()
	defer pool.Close()

	http.HandleFunc("/sensor", middleware.CheckAuth(handlers.GetSensorData(pool)))

	log.Fatal(http.ListenAndServe(serverPort, nil))
}
