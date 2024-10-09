package db

import (
	"auth-go-app/models"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite" //- here registers sql lite driver with database/sql package
)

var context *sql.DB

func Init() error {
	var err error
	context, err = sql.Open("sqlite", os.Getenv("SQLITE_DB_PATH"))
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	_, err = context.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			firstname TEXT,
			lastname TEXT,
			email TEXT UNIQUE,
			password TEXT,
			created_at DATETIME
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	return nil
}

func Close() {
	if context != nil {
		context.Close()
	}
}

func CheckIfUserExists(email string) bool {
	var count int
	err := context.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		fmt.Printf("Error checking if user exists: %v\n", err)
		return false
	}
	return count > 0
}

func Save(u *models.User) error {
	u.CreatedAt = time.Now()
	_, err := context.Exec(`
		INSERT INTO users (id, firstname, lastname, email, password, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt)

	if err != nil {
		return fmt.Errorf("error saving user: %v", err)
	}
	return nil
}
