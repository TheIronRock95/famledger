package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheIronRock95/famledger/internal/api"
	"github.com/TheIronRock95/famledger/internal/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect Supabase
	db.InitDatabase()

	// Router
	r := mux.NewRouter()

	// Register API routes
	api.RegisterRoutes(r)

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("FamLedger API is running 🚀"))
	})

	// Start server
	port := "8080"
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
