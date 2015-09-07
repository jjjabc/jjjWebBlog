package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/article"
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

func (this *ViewListController) GetCg() {
	page, err := strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}
	category:=this.GetString("cg")
	jas, err := article.GetPublishedArticlesByCategory(page, 10,category)
	if err != nil {
		this.Data["jas"] = nil
		return
	}
	this.Data["jas"] = jas
	this.Data["page"] = page
	this.TplNames = "viewlist.html"
}
func (this *ViewListController) GetCgJson(){
		page, err := strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}
	category:=this.GetString("cg")
	jas, err := article.GetPublishedArticlesByCategory(page, 10,category)
	if err != nil {
		this.Ctx.WriteString("{'error',"+err.Error()+"}")
		return
	}
	buf,err:=json.Marshal(jas)
	if err!=nil{
		this.Ctx.WriteString("{'error',"+err.Error()+"}")
		return
	}
	this.Ctx.WriteString(string(buf))
}