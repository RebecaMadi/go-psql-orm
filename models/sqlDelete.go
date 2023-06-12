package models


import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
)


func SqlDeleteDaily(db *sql.DB, d dbconfig.Daily){
	d = SqlSelectDaily(db, d)
	
	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id=$1", dbconfig.DaikyTableName)

	delete, err := db.Prepare(sqlStatement)
	CheckErr(err)

	result, err := delete.Exec(d.User)
	CheckErr(err)

	affect, err := result.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}

func SqlDelete(db *sql.DB, u dbconfig.Users){
	u = SqlSelect(db, u)
	
	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id=$1", dbconfig.TableName)

	delete, err := db.Prepare(sqlStatement)
	CheckErr(err)

	result, err := delete.Exec(u.Id)
	CheckErr(err)

	affect, err := result.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}