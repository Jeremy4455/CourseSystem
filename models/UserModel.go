package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       string `orm:"pk;size(20);column(id)"`
	Username string `orm:"size(20)"`
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
	user := &User{Id: id}
	err := o.Read(user)
	if err != nil {
		return false
	}
	o.Delete(user)
	return true
}

func ReviseUser(u *User, username, password string) bool {
	if u == nil {
		return false
	}
	if username != "" {
		u.Username = username
	}
	if password != "" {
		u.Password = password
	}
	return true
}
