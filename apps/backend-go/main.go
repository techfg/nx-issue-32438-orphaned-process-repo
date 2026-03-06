package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	host := envOrDefault("HOST", "127.0.0.1")
	port := envOrDefault("PORT", "3001")
	addr := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"Hello from backend-go"}`))
	})

	log.Printf("[ ready ] http://%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
