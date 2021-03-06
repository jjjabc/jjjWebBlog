package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/article"
	"strconv"
)

type AddartController struct {
	beego.Controller
}

func (this *AddartController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	ja := article.JJJarticle{}
	ja.Title = this.GetString("title")
	ja.Text = this.GetString("text")
	ja.Imgurl = this.GetString("imgurl")
	ja.IsPublished, _ = strconv.ParseBool(this.GetString("ispublish"))
	ja.Category = this.GetString("cg")
	if err := ja.AddArticle(); err != nil {
		beego.Info("add error")
		this.Ctx.WriteString("add error")
		return
	}
	if ja.IsPublished {
		if err := ja.Publish(); err != nil {
			beego.Info("publish error")
			this.Ctx.WriteString("publish error")
			return
		}
	}
	this.Ctx.WriteString("OK")
	return
}
