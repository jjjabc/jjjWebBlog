package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"strconv"
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
	red, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer red.Close()
	uid, err := redis.Int(red.Do("GET", "account:uname:"+username))
	beego.Info(uid)
	if err != nil {
		this.Ctx.Redirect(302, "./login")
		return
	} else if pwd, _ := redis.String(red.Do("GET", "account:password:"+strconv.Itoa(uid))); pwd != password {
		this.Ctx.Redirect(302, "./login")
		return
	}
	this.SetSession("username", username)
	this.Ctx.Redirect(302, "./chat")
}
