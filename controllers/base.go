package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Get() {
	c.TplName = "index.tpl"
}

func (c *BaseController) GetUserInfo() map[string]interface{} {
	name := c.GetSession("name").(string)
	userId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	userInfo := map[string]interface{}{
		"name":     name,
		"userId":   userId,
		"semester": semester,
	}
	return userInfo
}

func (c *BaseController) Err(err error) {
	c.Data["json"] = map[string]interface{}{"error": err.Error()}
}
func (c *BaseController) Sucess() {
	c.Data["json"] = map[string]interface{}{"message": "Success!"}
}
