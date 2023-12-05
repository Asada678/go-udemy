package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"

	_, err := Db.Exec(cmd, "taro", 20)

	if err != nil {
		log.Fatalln(err)
	}
}
