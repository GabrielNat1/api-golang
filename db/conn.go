package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
	configuration 
		 db
*/

/*
var (
	host = "go_db" string
	port = 5432 int
	user = "postgres" string
	password = "1234" int
	dbname = "postgres"

)*/

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//fmt.Println("connected to" + dbname)

	return db, nil
}
