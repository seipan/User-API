package db

import (
	"database/sql"
	"log"
)

func NewDriver() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=password dbname=testdb sslmode=disable")
	if err != nil {
		log.Println("db connect failed")
		panic(err)
	}
	log.Println("db connect success")

	return db
}
