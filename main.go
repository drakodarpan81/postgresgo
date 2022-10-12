package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "XXXX"
	password = "XXXX"
	dbname   = "XXXX"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user= %s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// instalar: go get github.com/lib/pq
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `INSERT INTO Employee(Name, EmpId) VALUES('David', 1)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynStmt := `INSERT INTO Employee (Name, EmpId) VALUES($1, $2)`
	_, e = db.Exec(insertDynStmt, "Laura", 2)
	CheckError(e)
	
}

func CheckError(err error)  {
	if err != nil {
		panic(err)
	}
}