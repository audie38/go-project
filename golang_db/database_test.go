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

func TestExecParameterizedSql(t *testing.T){
	db := GetConnection();
	defer db.Close()

	username := "audie"
	password := "milson"

	ctx := context.Background()
	query := "INSERT INTO USER(USERNAME, PASSWORD) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, username, password)

	if err != nil{
		panic(err)
	}

	insertedId, err := result.LastInsertId()

	if err != nil{
		panic(err)
	}

	fmt.Println("Inserted UserId : ", insertedId)
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

func TestQueryParameterizedSql(t *testing.T){
	db := GetConnection();
	defer db.Close()

	ctx := context.Background()
	username := "milson"
	password := "milson"

	query := "SELECT USERNAME FROM USER WHERE USERNAME = ? AND PASSWORD = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil{
		panic(err)
	}

	for rows.Next(){
		var(
			userName string
		)

		err := rows.Scan(&userName)
		if err != nil{
			panic(err)
		}

		fmt.Println(userName, "Login Success")
	}

	defer rows.Close()
}

func TestPrepareStatement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	qry := "INSERT INTO USER(USERNAME, PASSWORD) VALUES(?, ?)"

	stmnt, err := db.PrepareContext(ctx, qry)
	if err != nil{
		panic(err)
	}

	defer stmnt.Close()

	username := "test123"
	password := "password"

	result, err := stmnt.ExecContext(ctx, username, password)
	if err != nil{
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("Inserted Id: ", id)
}

func InsertUserTrx(db *sql.DB, ctx context.Context, userName string, password string){
	// Trx
	qry := "INSERT INTO USER(USERNAME, PASSWORD) VALUES(?, ?)"

	stmnt, err := db.PrepareContext(ctx, qry)
	if err != nil{
		panic(err)
	}

	defer stmnt.Close()

	result, err := stmnt.ExecContext(ctx, userName, password)
	if err != nil{
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("Inserted Id: ", id)
	// End Trx
}

func TestTransaction(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil{
		panic(err)
	}

	qry := "INSERT INTO USER(USERNAME, PASSWORD) VALUES(?, ?)"
	username := "admin"
	password := "admin"
	result, err := tx.ExecContext(ctx, qry, username, password)

	if err != nil{
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("Inserted Id: ", id)

	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil{
		panic(err)
	}
}