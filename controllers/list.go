package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type ListController struct {
	beego.Controller
}

func (this *ListController) Get() {
	jas, err := article.GetPublishedArticles(1, 10)
	if err != nil {
		beego.Info("error!")
		this.Data["msg"] = "error:" + err.Error()
	}
	beego.Info("jas len:" + strconv.Itoa(len(jas)))
	this.Data["jas"] = jas
	this.TplNames = "list.tpl"
}
