package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/beego")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	Database = db
}
