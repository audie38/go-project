package golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB{
	connString := "root:@tcp(localhost:3306)/golang_db?parseTime=true"
	db, err := sql.Open("mysql", connString)

	if err != nil{
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestExecSql(t *testing.T){
	db := GetConnection();
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO CUST(CUST_NAME, EMAIL, BALANCE, RATING, BIRTH_DT, MARRIED) VALUES('Milson', 'test1@localhost.com', 5000000, 5, '2023-05-05', false)"
	_, err := db.ExecContext(ctx, query)

	if err != nil{
		panic(err)
	}
}

func TestQuerySql(t *testing.T){
	db := GetConnection();
	defer db.Close()

	ctx := context.Background()
	query := "SELECT CUST_ID, CUST_NAME, EMAIL, BALANCE, RATING, BIRTH_DT, MARRIED, CREATED_AT FROM CUST"
	rows, err := db.QueryContext(ctx, query)

	if err != nil{
		panic(err)
	}

	for rows.Next(){
		var(
			id int
			name string
			email sql.NullString
			balance, rating float64
			birthDt sql.NullTime
			createdAt time.Time
			married bool
		)

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDt, &married, &createdAt)
		if err != nil{
			panic(err)
		}

		fmt.Println(id, name, email, balance, rating, birthDt, createdAt, married)
	}

	defer rows.Close()
}