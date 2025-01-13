package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

// quando o programa verifica se o ctx é done?
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Processing took more than 5 seconds")
		w.Write([]byte("Request sucessfully processed"))
	case <-ctx.Done():
		log.Println("Request canceled by the user")
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
	}
}
