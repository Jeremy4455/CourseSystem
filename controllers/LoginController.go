package controllers

import (
	"CourseSystem/models"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	id := c.GetString("username")
	password := c.GetString("password")
	user := models.Login(id, password)
	if user == nil {
		c.Redirect("/login", 302)
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
