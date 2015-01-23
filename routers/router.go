package routers

import (
	"github.com/astaxie/beego"
	"jjjBlog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/addart", &controllers.AddartController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/login", &controllers.LoginController{})
}
