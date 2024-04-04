package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id      string
	name    string
	sex     string
	college string
	major   string
}
type Teacher struct {
	id      string
	name    string
	sex     string
	college string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8&loc=Local")
	// register model
	orm.RegisterModel(new(Student))
	// create table
	orm.RunSyncdb("default", false, true)
}
