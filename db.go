package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db = getDb()

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nameOf_db")

	checkError(err, "Could not open db")

	return db
}
