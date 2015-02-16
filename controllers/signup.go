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
		c.Ctx.WriteString("SigupUser err:" + err.Error())
		return
	}
	c.Ctx.WriteString("OK")
}
