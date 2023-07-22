package main

import (
	"database/sql"
	"fmt"

	// Import the SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// Helper function to check if a table exists in the database
func isTableExists(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
