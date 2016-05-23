package main

import (
	"github.com/astaxie/beego"
)

type IndexCtr0 struct {
	beego.Controller
}

func (this *IndexCtr0) Get() {
	this.Ctx.WriteString("Hello Wolrd!")
}

func init() {
	beego.Router("/", &IndexCtr0{})
	beego.Run()
}
