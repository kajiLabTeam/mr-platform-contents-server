package model

import (
	"database/sql"
	"log"

	"github.com/kajiLabTeam/mr-platform-digital-twin-server/lib"
)

var db = connectDB()

func connectDB() *sql.DB {
	db, err := lib.SqlConnect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
