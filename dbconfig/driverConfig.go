package dbconfig

import (
	"fmt"
)

type Users struct{
	Id int
	Name string
	Email string
	Password string
	UserType int
}

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "Za932529"

const DbName = "agendapopam"

const TableName = "users"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
