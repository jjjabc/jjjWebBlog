package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/jjjabc/jjjWebBlog/orm"
	_ "github.com/jjjabc/jjjWebBlog/routers"
	"github.com/jjjabc/jjjWebBlog/user"
	"os"
)

func main() {
	IP := os.Getenv("REDIS_PORT_6379_TCP_ADDR")
	PORT := os.Getenv("REDIS_PORT_6379_TCP_PORT")
	PWD := os.Getenv("REDIS_PASSWORD")
	if IP == "" {
		IP = "127.0.0.1"
	}
	if PORT == "" {
		PORT = "6379"
	}
	if err := orm.Init(IP, PORT, PWD); err != nil {
		fmt.Println("orm init error!" + err.Error())
		return
	}
	exists, err := redis.Bool(orm.Red.Do("EXISTS", "account:count"))
	if err != nil {
		fmt.Println("orm init error!	" + err.Error())
		return
	}
	if !exists {
		admin := user.JJJuser{Name: "admin", Description: "初始管理员", NickName: "管理员"}
		err = admin.SigupUser("123456")
	}
	beego.SessionOn = true
	beego.Run()
}
