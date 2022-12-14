package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDriver() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=hoge dbname=hoge sslmode=disable")
	if err != nil {
		log.Println("db connect failed")
		panic(err)
	}
	log.Println("db connect success")

	err = db.Ping()
	if err != nil {
		log.Println("db connect failed")
		log.Println(err)
	}

	return db
}
