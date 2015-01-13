package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type JSONmsg struct {
	Num  int
	Info string
}
type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	c.TplNames = "signup.tpl"
}
func (c *SignupController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	red, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return
	}
	beego.Info("account:uname:" + username)
	if _, err := redis.String(red.Do("GET", "account:uname:"+username)); err == nil {
		msg := JSONmsg{
			Num:  1,
			Info: "Test Message!!!",
		}
		buf, _ := json.Marshal(msg)
		beego.Info("ok:")
		c.Ctx.ResponseWriter.Write(buf)
		return
	}
	uid, _ := redis.Int(red.Do("INCR", "account:count"))
	red.Do("SET", "account:uname:"+username, uid)
	red.Do("SET", "account:password:"+strconv.Itoa(uid), password)
	defer red.Close()
	c.Ctx.Redirect(302, "./login")
}
