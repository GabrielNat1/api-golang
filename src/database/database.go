package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Database configuration
const dbPath = "./database.db"

func InitDB() (*sql.DB, error) {
	// Create database directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	// Open SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Run migrations
	if err = runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	migrations := []string{
		"001_create_users.sql",
		"002_create_products.sql",
	}

	for _, migration := range migrations {
		content, err := os.ReadFile(filepath.Join("database/migrations", migration))
		if err != nil {
			log.Printf("Error reading migration %s: %v", migration, err)
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			log.Printf("Error executing migration %s: %v", migration, err)
			return err
		}
	}

	return nil
}
