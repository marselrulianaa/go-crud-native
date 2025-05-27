package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
}
