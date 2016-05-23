package controllers

import (
	"github.com/astaxie/beego"
)

type IndexCtr struct {
	beego.Controller
}

func (c *IndexCtr) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.tpl"
}
