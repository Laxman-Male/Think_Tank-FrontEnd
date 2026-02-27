package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"emergency-backend/db"
)

// Credentials represents login request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login checks credentials against the database and returns a JWT
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Example query - adjust columns/names to your schema
	var id uuid.UUID
	var storedHash string
	err := db.Conn.QueryRow("SELECT id, password_hash FROM users WHERE username=$1", creds.Username).Scan(&id, &storedHash)
	if err == sql.ErrNoRows {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// TODO: compare password with storedHash using bcrypt/scrypt etc.
	if creds.Password != storedHash {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := createToken(id.String())
	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func createToken(userID string) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	if len(signingKey) == 0 {
		return "", jwt.ErrTokenMalformed
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
