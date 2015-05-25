package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	jas, err := article.GetPublishedArticlesByCategory(1, 10, "top")
	if err != nil {
		beego.Info("error!")
		//		this.Data["msg"] = "error:" + err.Error()
	}
	//	beego.Info("IP:" + this.Ctx.Request.RemoteAddr + "	Host:" + this.Ctx.Request.Host)
	this.Data["jas"] = jas

	this.TplNames = "index.html"
}

func (this *MainController) GetTel() {
	jas, err := article.GetPublishedArticlesByCategory(1, 10, "tel")
	if err != nil {
		beego.Info("error!")
		//		this.Data["msg"] = "error:" + err.Error()
	}
	//	beego.Info("IP:" + this.Ctx.Request.RemoteAddr + "	Host:" + this.Ctx.Request.Host)
	this.Data["jas"] = jas

	this.TplNames = "index.html"
}
