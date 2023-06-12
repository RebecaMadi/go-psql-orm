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

type Daily struct{
	Hp_In int
	H_In int
	Hp_Out int
	H_Out int
	Ps string
	User Users
}

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "Za932529"

const DbName = "agendapopam"

const DaikyTableName = ""

const TableName = "users"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
