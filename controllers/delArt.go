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
	artId, _ := strconv.Atoi(this.GetString("artId"))
	ja := article.GetArticle(artId)
	ja.DelArticle()
	this.Ctx.WriteString("OK")
	return
}
