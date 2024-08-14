package store

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func NewSQLStorage(DBSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DBSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}