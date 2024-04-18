package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	web.Controller
}

func (c *LogoutController) Get() {
	c.TplName = "logout.tpl"
	err := c.DestroySession()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("ok")
}
