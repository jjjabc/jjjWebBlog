package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	jas, err := article.GetPublishedArticles(1, 10)
	if err != nil {
		beego.Info("error!")
		this.Data["msg"] = "error:" + err.Error()
	}
	beego.Info("IP:" + this.Ctx.Request.RemoteAddr + "	Host:" + this.Ctx.Request.Host)
	this.Data["jas"] = jas

	this.TplNames = "index.tpl"
}
