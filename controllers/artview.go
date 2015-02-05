package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type ArtViewController struct {
	beego.Controller
}

func (this *ArtViewController) Get() {

	artId, err := strconv.Atoi(this.GetString("artid"))
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	ja := article.GetArticle(artId)
	if ja == nil {
		this.Ctx.WriteString("GetArticle error")
		return
	} else if !ja.IsPublished {
		this.Ctx.WriteString("Article is not published!")
		return
	}
	this.Data["ja"] = ja
	this.TplNames = "artview.tpl"
}
