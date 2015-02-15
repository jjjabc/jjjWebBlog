package routers

import (
	"github.com/astaxie/beego"
	"jjjBlog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/addArt", &controllers.AddartController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/artList", &controllers.ArtListController{})
	beego.Router("/publishArt", &controllers.PublishArtController{})
	beego.Router("/delArt", &controllers.DelArtController{})
	beego.Router("/updataArt", &controllers.UpdataArtController{})
	beego.Router("/upload", &controllers.UploadFileController{})
	beego.Router("/viewart", &controllers.ArtViewController{})
	beego.Router("/viewlist", &controllers.ViewListController{})
}
