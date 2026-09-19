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

	addr := ":" + port
	log.Printf("server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello from Mondoo Engineer!")
}