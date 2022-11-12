package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	conn := "user=postgres dbname=alura_store password=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}
