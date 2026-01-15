package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	log.Println("in echoString")
	fmt.Fprintf(w, "hello")
}

func main() {
	fmt.Println("Hi there")
	http.HandleFunc("/", echoString)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
