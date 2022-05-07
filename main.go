package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "deepak"
	dbname   = "postgres"
)

func main() {
	// connection string
	db := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	conn, err := gorm.Open("postgres", db)
	db1, err := sql.Open("postgres", db)
	if err != nil {
		fmt.Print(err)
	}
	// close database
	defer conn.Close()

	// check db
	err = conn.DB().Ping()
	CheckError(err)

	fmt.Println("Connected!")

	insertStmt := `insert into "LoginApi"("name", "email", "country","phone","password") values('John', 'John@gmail.com','brazil','21312','John')`
	_, e := db1.Exec(insertStmt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
