package connectiondb

import "database/sql"

func OpenDBPQL() *sql.DB {
	db, err := sql.Open("postgress", "postgres://postgres:postgres@localhost/graphql_db?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
