package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

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

	// 1. Define your existing routes
	r.HandleFunc("/", handlers.Homepage).Methods("GET")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	// Add your protected handlers here

	// 2. Add the CORS configuration here
	// This solves the 'strict-origin-when-cross-origin' error
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Your Angular app
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true, // Set to false in production
	})

	// 3. Wrap your router 'r' with the CORS handler 'c'
	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Starting server on :%s\n", port)

	// 4. IMPORTANT: Use 'handler' here, NOT 'r'
	log.Fatal(http.ListenAndServe(":"+port, handler))

	// r := mux.NewRouter()

	// r.HandleFunc("/", handlers.Homepage).Methods("GET")
	// r.HandleFunc("/login", handlers.Login).Methods("POST")

	// protected := r.PathPrefix("/api").Subrouter()
	// protected.Use(middleware.JWTMiddleware)
	// protected.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("authenticated hello"))
	// })

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8000"
	// }

	// fmt.Printf("Starting server on :%s\n", port)
	// log.Fatal(http.ListenAndServe(":"+port, r))
}
