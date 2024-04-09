package models

type Admin struct {
	Id       int    `orm:"pk;auto;column(id)"`
	Username string `orm:"size(16);unique"`
	Password string `orm:"size(32)"`
}

func (this *Admin) Login() {
	return
}
