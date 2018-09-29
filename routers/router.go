package routers

import (
	"beego3/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/admin", &controllers.UserController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/index", &controllers.IndexController{})
}
