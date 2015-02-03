package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type DelArtController struct {
	beego.Controller
}

func (this *DelArtController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	artId, err := strconv.Atoi(this.GetString("artId"))
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	ja := article.GetArticle(artId)
	if ja == nil {
		beego.Info("get err")
		this.Ctx.WriteString("Getarticle error" + strconv.Itoa(artId))
		return
	}
	ja.DelArticle()
	this.Ctx.WriteString("OK")
	return
}
