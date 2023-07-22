package main

import (
	"database/sql"
	"fmt"
	"ovi/model"
)

func main() {
	u := model.NewUltraman()
	u.SetName("อุลตร้าแมน")
	u.SetTool("อุลตร้าแมนเทิล")
	u.SetSpecialAttack("ลำแสงสเปเซี่ยม")
	u.SetTransformKit("เบต้าแคปซูล")

	fmt.Println(u)

	fmt.Println(u.Json())

	// Init Table
	// Connect to the SQLite database (creates a new file if it doesn't exist)
	db, err := sql.Open("sqlite3", "ultraman.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Check if the table 'ultramen' exists
	exists, err := isTableExists(db, "ultramen")

	// ถ้า err != nil -> 1+2=3 = condition
	if err != nil {
		fmt.Println("Error checking table existence:", err)
		return
	}

	// If the table does not exist, create it
	if !exists {
		createTableSQL := `
			CREATE TABLE ultramen (
				name TEXT,
				tool TEXT,
				specialAttack TEXT,
				transformKit TEXT
			);
		`
		_, err = db.Exec(createTableSQL)
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}

		fmt.Println("Table 'ultramen' created successfully.")
	} else {
		fmt.Println("Table 'ultramen' already exists.")
	}

}
