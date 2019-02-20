package routers

import (
	"github.com/TimeBye/amap/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/around", &controllers.AroundController{})
}
