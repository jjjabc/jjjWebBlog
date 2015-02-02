package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type ArtListController struct {
	beego.Controller
}

func (this *ArtListController) Get() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	jas, err := article.GetAllArticles()
	if err != nil {
		beego.Info("error!")
		this.Ctx.WriteString("Get articles error!")
	}
	beego.Info("jas len:" + strconv.Itoa(len(jas)))
	this.Data["jas"] = jas
	this.TplNames = "artList.tpl"
}
