package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/user"
)

type JSONmsg struct {
	Num  int
	Info string
}
type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	v := c.GetSession("username")
	if v == nil {
		c.Ctx.WriteString("Not login！")
		return
	}
	c.TplNames = "signup.tpl"
}
func (c *SignupController) Post() {
	v := c.GetSession("username")
	if v == nil {
		c.Ctx.WriteString("Not login！")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")

	beego.Info("account:uname:" + username)

	ju := user.JJJuser{
		Name: username,
	}
	err := ju.SigupUser(password)
	if err != nil {
		beego.Info("SigupUser err:" + err.Error())
		c.Ctx.Redirect(302, "./login")
		return
	}
	c.Ctx.Redirect(302, "./login")
}
