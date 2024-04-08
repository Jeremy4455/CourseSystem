package models

type Admin struct {
	Id       int    `orm:"column(Id)"`
	Username string `orm:"size(16);unique;column(Username)"`
	Password string `orm:"size(32);column(Password)"`
}

func (this *Admin) Login() {
	return
}
