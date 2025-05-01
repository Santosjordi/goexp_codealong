package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Simulated panic")
	})

	log.Print("Server is running on port 3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
