package routers

import (
	"helloworld/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/home/", &controllers.IndexCtr{} )
}
