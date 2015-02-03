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
	artId, err := strconv.Atoi(this.GetString("artId"))
	if err != nil {
		this.Ctx.WriteString("Get artId=" + this.GetString("artId") + " error:" + err.Error())
		return
	}
	ja := article.GetArticle(artId)
	if ja == nil {
		this.Ctx.WriteString("Get artcle error")
		return
	}
	if this.GetString("Action") == "publish" {
		err = ja.Publish()
	} else {
		err = ja.UnPublish()
	}
	if err != nil {
		beego.Info("publish/unpublish error")
		this.Ctx.WriteString("Pubilsh/Unpublish article error:" + err.Error())
	} else {
		this.Ctx.WriteString("OK")
	}
	return
}
