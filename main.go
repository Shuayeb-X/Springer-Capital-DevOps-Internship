package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Root endpoint handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"application": "Springer Capital Demo App",
		"status":      "running",
	})
}

// Health endpoint handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

// Register application routes
func newHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)

	return mux
}

func main() {
	log.Println("Springer Demo App starting on port 8080")

	err := http.ListenAndServe(":8080", newHandler())

	if err != nil {
		log.Fatal(err)
	}
}
