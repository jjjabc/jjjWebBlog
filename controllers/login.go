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
		this.Ctx.WriteString("成功")
		return
	} else {
		this.Ctx.WriteString("错误")
		return
	}

}
