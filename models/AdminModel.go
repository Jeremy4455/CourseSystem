package models

type Admin struct {
	Id       string `orm:"pk;size(16);auto;column(id)"`
	Username string `orm:"size(16);unique"`
	Password string `orm:"size(32)"`
}
