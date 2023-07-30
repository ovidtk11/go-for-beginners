package main

import (
	"database/sql"
	"fmt"
	"ovi/database"
	"ovi/model"
)

var conn *sql.DB

func main() {

	// Connect to the SQLite database (creates a new file if it doesn't exist)
	openConnectDB()

	// Init Table
	createTable()

	// create untraman data for insert to database
	data := createUltranManData()

	// Insert sample data into the 'ultramen' table
	insertUlTraMan(data)

}

func openConnectDB() {
	db, err := sql.Open("sqlite3", "ultraman.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	conn = db
}

func createTable() {
	// Check if the table 'ultramen' exists
	exists, err := database.IsTableExists(conn, "ultramen")

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
		_, err = conn.Exec(createTableSQL)
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}

		fmt.Println("Table 'ultramen' created successfully.")
	} else {
		fmt.Println("Table 'ultramen' already exists.")
	}
}

// slice -> กล่องที่เราต้องการเก็บของ หลายๆชิ้น เช่น object (ex. model.IUltraman)  เราต้องการเก็บ taro, leo ไว้ในกล่อง(ด้วยกัน)
// เราก็จะสร้าง slice ขึ้นมาเพื่อเก็บของ (object หรือ type นั้นๆ) * slice สามารถบรรจุของได้แค่ type ที่มันประกาศไว้เท่านั้นนะ *
// data => [slice]<<Type>> -> []model.IUltraman -> [{model.IUltraman},{model.IUltraman}] -> [taro, leo]

func createUltranManData() (data []model.IUltraman) {
	ultramenData := make([]model.IUltraman, 0)

	taro := model.NewUltraman()
	taro.SetName("Ultaman Taro")
	taro.SetTool("Florium Beam")
	taro.SetSpecialAttack("Storium Ray")
	taro.SetTransformKit("Bracelet")

	leo := model.NewUltraman()
	leo.SetName("Ultraman Leo")
	leo.SetTool("Leo Nunchaku")
	leo.SetSpecialAttack("Leo Kick")
	leo.SetTransformKit("Leo Ring")

	ultramenData = append(ultramenData, taro, leo)

	return ultramenData
}

func insertUlTraMan(data []model.IUltraman) {
	for _, u := range data {
		insertSQL := `
			INSERT INTO ultramen (name, tool, specialAttack, transformKit)
			VALUES (?, ?, ?, ?);
		`

		_, err := conn.Exec(insertSQL, u.GetName(), u.GetTool(), u.GetSpecialAttack(), u.GetTransformKit())
		if err != nil {
			fmt.Println("Error inserting data:", err)
			return
		}
	}

	fmt.Println("Data inserted into 'ultramen' table successfully.")
}
