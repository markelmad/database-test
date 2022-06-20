package helper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:markel258@tcp(127.0.0.1:3306)/zestyiotest")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
