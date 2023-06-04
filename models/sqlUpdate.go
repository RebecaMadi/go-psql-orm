package models

import (
	"database/sql"
	"fmt"

	"github.com/go-psql-orm/dbconfig"
)

func SqlUpdate(db *sql.DB, u dbconfig.Users, field string, content string){
	u = SqlSelect(db, u)
	
	sqlStatement := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE id = %d", dbconfig.TableName, field, u.Id)

	update, err := db.Prepare(sqlStatement)
	CheckErr(err)

	result, err := update.Exec(content)
	CheckErr(err)

	affect, err := result.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}