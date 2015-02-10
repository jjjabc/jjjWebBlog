package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/user"
)

type JSONmsg struct {
	Num  int
	Info string
}
type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	key := c.GetString("key")
	if key != "jjjabc" {
		c.Ctx.WriteString("No login")
		return
	}
	c.TplNames = "signup.tpl"
}
func (c *SignupController) Post() {
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
