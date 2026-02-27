package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"emergency-backend/db"
	"emergency-backend/handlers"
	"emergency-backend/middleware"
)

func main() {
	// load .env if present
	_ = godotenv.Load()

	if err := db.Connect(); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Homepage).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("authenticated hello"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Starting server on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
