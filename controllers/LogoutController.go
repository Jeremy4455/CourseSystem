package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	web.Controller
}

func (c *LogoutController) Get() {
	c.TplName = "logout.tpl"
	err := c.DestroySession()
	if err != nil {
		return
	}
}
