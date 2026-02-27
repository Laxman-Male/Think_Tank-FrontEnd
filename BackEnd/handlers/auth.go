package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"emergency-backend/db"
)

// Credentials represents login request body
type Credentials struct {
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	IsNewUser bool   `json:"isNewUser"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// 1. Check if Mobile Number already exists
	var exists int
	err := db.Conn.QueryRow("SELECT COUNT(*) FROM users WHERE mobile = ?", creds.Mobile).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists > 0 {
		http.Error(w, "Mobile number already registered", http.StatusConflict)
		return
	}

	// 2. Hash Password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	// 3. Insert into MySQL (Email is optional, handled automatically by Go/MySQL if empty)
	// We include 'token' as an empty string initially or handle it in the next step
	query := "INSERT INTO users (email, mobile, password_hash, role, token) VALUES (?, ?, ?, ?, ?)"

	// 4. Generate Token FIRST so we can save it in the same INSERT or a subsequent UPDATE
	// Since we don't have the ID yet, we'll INSERT then UPDATE (most stable way)
	result, err := db.Conn.Exec(query, creds.Email, creds.Mobile, string(passwordHash), creds.Role, "")
	if err != nil {
		http.Error(w, "Could not register user", http.StatusInternalServerError)
		return
	}

	// 5. Get the new ID to generate the JWT
	newID, _ := result.LastInsertId()
	tokenString, err := createToken(fmt.Sprintf("%d", newID))
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	// 6. UPDATE the record with the actual token
	_, err = db.Conn.Exec("UPDATE users SET token = ? WHERE id = ?", tokenString, newID)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	// 7. Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":   tokenString,
		"role":    creds.Role,
		"message": "Registration successful",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// 1. Prepare variables to hold DB results
	var id int
	var storedHash string
	var userRole string
	var mobile string
	var email sql.NullString // Use NullString because email is optional (can be NULL)

	// 2. Fetch data based on Mobile AND Role
	query := "SELECT id, password_hash, role, mobile, email FROM users WHERE mobile = ? AND role = ?"
	err := db.Conn.QueryRow(query, creds.Mobile, creds.Role).Scan(&id, &storedHash, &userRole, &mobile, &email)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid mobile number or role", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// 3. Compare Bcrypt Hash
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// 4. Create Token
	tokenString, err := createToken(fmt.Sprintf("%d", id))
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	// 5. Update the token in the DB (for session management)
	_, _ = db.Conn.Exec("UPDATE users SET token = ? WHERE id = ?", tokenString, id)

	// 6. Build the response map
	response := map[string]interface{}{
		"token":  tokenString,
		"id":     id,
		"mobile": mobile,
		"role":   userRole,
	}

	// Only add email to response if it's not NULL in the DB
	if email.Valid {
		response["email"] = email.String
	} else {
		response["email"] = ""
	}

	// 7. Send Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	// Fallback logic for local development if .env fails to load
	if secret == "" {
		fmt.Println("WARNING: JWT_SECRET not found in .env, using default secret")
		secret = "emergency_default_secret_key"
	}

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
