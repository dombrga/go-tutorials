package sqldb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:sample@127.0.0.1/practice_sample?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}
