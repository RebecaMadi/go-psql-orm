package main

import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
    _ "github.com/go-psql-orm/models"

	_ "github.com/lib/pq"
)

var db *sql.DB
var user dbconfig.Users;
var err error

func main() {

    name := "Nome"
    email := "nome@gmail.com"
    password := "123456789"
    usertype := 1;

    user.Name = name
    user.Email = email
    user.Password = password
    user.UserType = usertype

	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

}
