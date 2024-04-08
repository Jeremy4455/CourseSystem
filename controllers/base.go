package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Get() {
	this.TplName = "index.tpl"
}
