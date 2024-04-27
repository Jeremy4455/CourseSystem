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
	c.SetSession("name", user.Username)

	switch role {
	case "admin":
		c.SetSession("userId", "0"+id)
		c.Redirect("/admin", 302)
	case "student":
		c.SetSession("userId", "1"+id)
		c.Redirect("/student", 302)
	case "teacher":
		c.SetSession("userId", "2"+id)
		c.Redirect("/teacher", 302)
	}
}
