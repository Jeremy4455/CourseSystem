package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       string `orm:"pk;size(20);auto;column(id)"`
	Username string `orm:"unique;size(20)"`
	Password string `orm:"size(32);column(Password)"`
	Role     string `orm:"size(32);column(Role)"`
}

func GetUser(id string) bool {
	return true
}

func AddUser(id, username, password, role string) bool {
	o := orm.NewOrm()

	user := &User{Id: id, Username: username, Password: password, Role: role}
	o.Insert(user)
	return true
}

func DeleteUser(id string) bool {
	o := orm.NewOrm()

	user := o.Read(&User{Id: id})
	o.Delete(user)
	return true
}
