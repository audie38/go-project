package helper

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	connString := "root:@tcp(localhost:3306)/golang_db?parseTime=true"
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}