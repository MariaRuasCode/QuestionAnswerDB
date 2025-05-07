package migrate

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func RunMigrations() error {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=qa_backend sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer db.Close()

	// Adjusted path for Windows & relative location
	migrationSQL, err := os.ReadFile("../../migrations/001_init.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	fmt.Println("✅ Migration executed successfully.")
	return nil
}
