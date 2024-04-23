package models

import (
	"github.com/beego/beego/orm"
)

type User struct {
	Id       string `orm:"pk;size(20);column(id)"`
	Username string `orm:"size(20)"`
	Password string `orm:"size(32);column(Password)"`
	Role     string `orm:"size(32);column(Role)"`
}

func Login(id, password string) *User {
	user := &User{Id: id}
	o := orm.NewOrm()
	err := o.Read(user)
	if err != nil {
		return nil
	}
	if user.Password != password {
		return nil
	}
	return user
}

func GetUser(id string) bool {
	return true
}

func AddUser(id, username, password, role string) error {
	o := orm.NewOrm()

	user := &User{Id: id, Username: username, Password: password, Role: role}
	_, err := o.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)
	if err != nil {
		return err
	}

	_, err = o.Delete(user)
	if err != nil {
		return err
	}
	return nil
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
