package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type UpdataArtController struct {
	beego.Controller
}

func (this *UpdataArtController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	artId, _ := strconv.Atoi(this.GetString("id"))
	ja := article.GetArticle(artId)
	ja.Title = this.GetString("title")
	ja.Text = this.GetString("text")
	if err := ja.UpdataArticle(); err != nil {
		beego.Info("Updata error")
		this.Ctx.WriteString("updata error")
		return
	}

	this.Ctx.WriteString("OK")
}
