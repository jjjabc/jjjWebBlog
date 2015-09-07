package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/orm"
	_ "github.com/jjjabc/jjjWebBlog/routers"
)

func main() {
	if err := orm.InitPass("10.10.71.37", "50320","RPs0uYV1"); err != nil {
		fmt.Println("orm init error!")
		return
	}
	beego.SessionOn = true
	beego.Run()
}
