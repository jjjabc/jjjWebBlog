package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type PublishArtController struct {
	beego.Controller
}

func (this *PublishArtController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	artId, _ := strconv.Atoi(this.GetString("artId"))
	ja := article.GetArticle(artId)
	if this.GetString("Action") == "publish" {
	}
	if err := ja.Publish(); err != nil {
		beego.Info("publish error")
		this.Ctx.WriteString("Pubilsh article error!")

	} else {
		this.Ctx.WriteString("OK")
	}
	return
}
