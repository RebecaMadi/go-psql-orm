package models

import (
	"database/sql"
	"fmt"
	"github.com/go-psql-orm/dbconfig"
)

func SqlSelectDaily(db *sql.DB, d dbconfig.Daily) (dl dbconfig.Daily){

	st := fmt.Sprintf("SELECT id, name, email, password, user_type FROM %s WHERE email = '%s'", dbconfig.DaikyTableName, d.User.Email)
	
	sqlStatement, err := db.Query(st)
	CheckErr(err)
	for sqlStatement.Next() {
		err = sqlStatement.Scan(&d.Hp_In, &d.H_In, &d.Hp_Out, &d.H_Out, &d.Ps)
		CheckErr(err)
	}
	
	//fmt.Println(u.Id, u.Name, u.Email, u.Password, u.UserType)

	return d;
}


func SqlSelect(db *sql.DB, u dbconfig.Users) (usr dbconfig.Users){

	st := fmt.Sprintf("SELECT id, name, email, password, user_type FROM %s WHERE email = '%s'", dbconfig.TableName, u.Email)
	
	sqlStatement, err := db.Query(st)
	CheckErr(err)
	for sqlStatement.Next() {
		err = sqlStatement.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.UserType)
		CheckErr(err)
	}
	
	fmt.Println(u.Id, u.Name, u.Email, u.Password, u.UserType)

	return u;
}
