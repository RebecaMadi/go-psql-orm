package models

import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
)

func SqlInsertDaily(db *sql.DB, d dbconfig.Daily) {
	sqlStatement := fmt.Sprintf("INSERT INTO %s(name, email, password, user_type) VALUES($1, $2, $3, $4)", dbconfig.DaikyTableName)

	insert, err := db.Prepare(sqlStatement)
	CheckErr(err)

	result, err := insert.Exec(d.Hp_In, d.H_In, d.Hp_Out, d.H_Out, d.Ps, d.User)
	CheckErr(err)

	affect, err := result.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}

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
