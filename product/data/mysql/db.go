package mysql

import (
	"database/sql"
	"log"
)

var DbConnection *sql.DB

func SetupDB() {
	var err error
	DbConnection, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/inventorydb")

	if err != nil {
		log.Fatal(err)
	}
}
