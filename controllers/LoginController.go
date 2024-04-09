package controllers

import (
	"CourseSystem/models"
	"errors"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	o := orm.NewOrm()
	user := models.User{Username: username}

	err := o.Read(&user, "Username")
	if err != nil {
		this.Redirect("/login", 302)
		errors.New(
			"不存在该用户")
		return
	}
	if user.Password != password {
		this.Redirect("/login", 302)
		errors.New(
			"密码错误")
		return
	}
	role := user.Role
	switch role {
	case "admin":
		this.Redirect("/admin", 302)
	case "student":
		this.Redirect("/student", 302)
	case "teacher":
		this.Redirect("/teacher", 302)
	}
}
