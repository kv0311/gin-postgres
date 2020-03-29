package repo

import (
	"database/sql"
	"log"
)

func connect() (db *sql.DB, err error) {
	connStr := "user=khanhvinh dbname=postgres host=localhost password=khanhvinh1998 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
