package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/user"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"

}
func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	beego.Info(username + ":" + password)
	if user.CheckUser(username, password) {
		this.SetSession("username", username)
		this.Ctx.Redirect(302, "/admin")

		return
	} else {
		this.Ctx.WriteString("登陆错误")
		return
	}

}
