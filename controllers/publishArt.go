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
	var err error
	if this.GetString("Action") == "publish" {
		err = ja.Publish()
	} else {
		err = ja.UnPublish()
	}
	if err != nil {
		beego.Info("publish/unpublish error")
		this.Ctx.WriteString("Pubilsh/Unpublish article error!")
	} else {
		this.Ctx.WriteString("OK")
	}
	return
}
