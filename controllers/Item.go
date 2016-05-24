package controllers

import (
	"github.com/astaxie/beego"
)

type ItemCtr struct {
	beego.Controller
}

func (this *ItemCtr) Prepare(){
        // run after Init
}

func (this *ItemCtr) Get() {
	name := "HouLinwei"
	var id string
	id = this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.Data["name"] = name
	this.TplName = "item.tpl"
}

func (this *ItemCtr) Post() {
	// save a item to db.
}

func (this *ItemCtr) Finish() {
	// close db connection here.
}
