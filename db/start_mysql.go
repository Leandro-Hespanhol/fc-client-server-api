package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func StartMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/exchanges")
	if err != nil {
		panic(err)
	}

	fmt.Println("MySQL connected")
	createExchangeTable(db)
	return db
}

func createExchangeTable(db *sql.DB) {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS exchanges (id INT AUTO_INCREMENT PRIMARY KEY, bid DECIMAL(10, 4), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);")
	if err != nil {
		panic(err)
	}
	log.Println(result)

	log.Println("exchanges table created")
}
