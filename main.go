package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/orm"
	_ "github.com/jjjabc/jjjWebBlog/routers"
)

func main() {
	if err := orm.Init("127.0.0.1", "6379"); err != nil {
		fmt.Println("orm init error!")
		return
	}
	beego.SessionOn = true
	beego.Run()
}
