package models

type User struct {
	Id       int    `orm:"pk;auto;column(id)"`
	Username string `orm:"unique;size(20)"`
	Password string `orm:"size(32);column(Password)"`
	Role     string `orm:"size(32);column(Role)"`
}
