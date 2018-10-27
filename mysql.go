package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/parkir")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)

	return db
}
