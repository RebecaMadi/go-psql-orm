package main

import (
	"database/sql"
	"fmt"

	"./gopostigres/dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	/*sqlSelect()
	  sqlSelectID()
	  sqlInsert()
	  sqlUpdate()
	  sqlDelete()*/
}
