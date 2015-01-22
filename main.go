package main

import (
	"github.com/astaxie/beego"
	"jjjBlog/orm"
	_ "jjjBlog/routers"
)

func main() {
	if err := orm.Init("127.0.0.1", "6379"); err != nil {
		beego.Info("orm init error:")
	}
	beego.Run()
}
