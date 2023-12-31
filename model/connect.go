package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/prg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con = db
	return db

}
