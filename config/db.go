package config

import "database/sql"

func ConnectDB() *sql.DB {
	connection := "user=usernamepg dbname=db password=password123 host=postgresgo sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
