package controllers

import (
	"github.com/astaxie/beego"
)

type IndexCtr struct {
	beego.Controller
}

func (this *IndexCtr) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "home.tpl"
}
