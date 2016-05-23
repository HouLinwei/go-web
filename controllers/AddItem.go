package controllers

import (
	"github.com/astaxie/beego"
)

type AddItemCtr struct {
	beego.Controller
}

func (this *AddItemCtr) Get() {
	name := "HouLinwei"
	var id string
	id = this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.Data["name"] = name
	this.TplName = "item.tpl"
}
