package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DatabaseConnection = *dbConn()

func dbConn() *sql.DB {
	connSrt := "postgres://postgres:<<PASSWORD>>@localhost:32770/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connSrt)
	if err != nil {
		log.Fatal("Error DB connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	return db
}
