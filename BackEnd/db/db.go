package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

var Conn *sql.DB

// Connect opens a database connection using environment variables.
func Connect() error {
	// Build DSN from individual environment variables for flexibility
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || name == "" {
		return fmt.Errorf("database environment variables (DB_USER, DB_PASS, DB_HOST, DB_NAME) must be set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, name)

	var err error
	Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Verify connection
	if err := Conn.Ping(); err != nil {
		return err
	}

	return nil
}
