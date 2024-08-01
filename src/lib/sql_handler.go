package lib

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func SqlConnect() (*sql.DB, error) {
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_CONTENTS")

	db, err := sql.Open("postgres", "host=postgres port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
