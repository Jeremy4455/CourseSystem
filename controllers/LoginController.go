package controllers

import (
	"CourseSystem/models"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	o := orm.NewOrm()
	user := models.User{Username: username}

	err := o.Read(&user, "Username")
	if err != nil {
		c.Redirect("/login", 302)
		//errors.New("不存在该用户")
		return
	}
	if user.Password != password {
		c.Redirect("/login", 302)
		//errors.New("密码错误")
		return
	}
	role := user.Role
	switch role {
	case "admin":
		c.Redirect("/admin", 302)
	case "student":
		c.Redirect("/student", 302)
	case "teacher":
		c.Redirect("/teacher", 302)
	}
}
