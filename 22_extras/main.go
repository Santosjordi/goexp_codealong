package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3000"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		// Simulate a panic
		mockPanic()
	})

	go func() {
		fmt.Println("Server is running on port 3000")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down server:", err)
	}

	fmt.Println("Server stopped")

}

func mockPanic() {
	// This function is intentionally left empty to simulate a panic.
	// In a real-world scenario, this would be replaced with actual code that could cause a panic.
	panic("This is a simulated panic")
}
