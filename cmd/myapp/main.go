package main

import (
	"fmt"
	"net/http"
	"os"

	"myapp/internal/compute"

	"github.com/google/uuid"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := 25
		fib := compute.Fib(n)
		id := uuid.New() // external dependency
		fmt.Fprintf(w, "hello from myapp! fib(%d)=%d uuid=%s\n", n, fib, id)
	})

	// Read PORT from environment, fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port " + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
