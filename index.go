package main

import (
    "github.com/astaxie/beego"
)

type IndexCtr struct{
    beego.Controller
}

func (this *IndexCtr) Get(){
    this.Ctx.WriteString("Hello Wolrd!")
}

func main(){
    beego.Router("/", &IndexCtr{})
    beego.Run()
}
