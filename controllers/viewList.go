package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type ViewListController struct {
	beego.Controller
}

func (this *ViewListController) Get() {
	page, err := strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}
	jas, err := article.GetPublishedArticles(page, 10)
	if err != nil {
		this.Data["jas"] = nil
		return
	}
	this.Data["jas"] = jas
	this.Data["page"] = page
	this.TplNames = "viewlist.html"
}
