package routers

import (
	"tom/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/home/", &controllers.IndexCtr{} )
    beego.Router("/dpl/?:id", &controllers.ItemCtr{} )
}
