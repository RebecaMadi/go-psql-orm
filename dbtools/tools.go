package dbtools
import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
    _ "github.com/go-psql-orm/models"

	_ "github.com/lib/pq"
)

var db *sql.DB
//var user dbconfig.Users;
var err error

func DbOpen() (b *sql.DB){
	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
	return db
}

func DbClose(b *sql.DB){
	b.Close()
}
/*
func ApiORM(email string, password string) {
    user.Email = email
    user.Password = password

	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

}
*/