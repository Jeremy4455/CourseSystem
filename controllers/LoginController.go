package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	// 假设这里有一个验证用户名密码的逻辑
	if username == "admin" && password == "password123" {
		//// 登录成功，将用户信息存储到 Session 中
		//c.SetSession("username", username)

		// 重定向到登录成功后的页面
		c.Redirect("/", 302)
		return
	}

	// 登录失败，重新回到登录页面
	c.Redirect("/login", 302)
}
