package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	addr := ":" + port
	log.Printf("server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "Hello from Mondoo Engineer!"); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}