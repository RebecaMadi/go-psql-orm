package models

import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
)

func SqlInsert(db *sql.DB, u dbconfig.Users) {
	sqlStatement := fmt.Sprintf("INSERT INTO %s(name, email, password, user_type) VALUES($1, $2, $3, $4)", dbconfig.TableName)

	insert, err := db.Prepare(sqlStatement)
	CheckErr(err)

	result, err := insert.Exec(u.Name, u.Email, u.Password, u.UserType)
	CheckErr(err)

	affect, err := result.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}
