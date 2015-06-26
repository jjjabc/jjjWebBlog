package controllers

import (
	"encoding/json"
	"jjjBlog/article"

	"github.com/astaxie/beego"
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
func (this *MainController) GetTelJson() {
	jas, err := article.GetPublishedArticlesByCategory(1, 10, "tel")
	if err != nil {
		beego.Info("error!")
		return
		//		this.Data["msg"] = "error:" + err.Error()
	}
	//	beego.Info("IP:" + this.Ctx.Request.RemoteAddr + "	Host:" + this.Ctx.Request.Host)
	buf, err := json.Marshal(jas)
	if err != nil {
		this.Ctx.WriteString("{'error'," + err.Error() + "}")
		return
	}
	this.Ctx.WriteString(string(buf))
}
func (this *MainController) GetJson() {
	jas, err := article.GetPublishedArticlesByCategory(1, 10, "top")
	if err != nil {
		beego.Info("error!")
		return
		//		this.Data["msg"] = "error:" + err.Error()
	}
	//	beego.Info("IP:" + this.Ctx.Request.RemoteAddr + "	Host:" + this.Ctx.Request.Host)
	buf, err := json.Marshal(jas)
	if err != nil {
		this.Ctx.WriteString("{'error'," + err.Error() + "}")
		return
	}
	this.Ctx.WriteString(string(buf))
}
