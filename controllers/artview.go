package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
	"strconv"
)

type ArtViewController struct {
	beego.Controller
}

type artTitle struct {
	Id    int
	Title string
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
		v := this.GetSession("username")
		if v == nil {
			this.Ctx.WriteString("Article is not published!")
			return
		}
		//管理员可以查看未发布文章，匿名就报错
	}
	preId, nextId := ja.GetRoundId()
	preview := article.GetArticle(preId)
	next := article.GetArticle(nextId)
	/*	var next *artTitle
		nextId, err := strconv.Atoi(this.GetString("nextid"))
		if err == nil {
			jart := article.GetArticle(nextId)
			if jart == nil {
				next = nil
			} else {
				next = new(artTitle)
				next.Id = nextId
				next.Title = jart.Title
			}
		} else {
			preId, nextId = ja.GetRoundId()
		}
		var preview *artTitle
		preId, err := strconv.Atoi(this.GetString("preid"))
		if err == nil {
			jart := article.GetArticle(preId)
			if jart == nil {
				preview = nil
			} else {
				preview = new(artTitle)
				preview.Id = preId
				preview.Title = jart.Title
			}
		} else {
			preId, nextId = ja.GetRoundId()
		}
	*/
	this.Data["ja"] = ja
	this.Data["next"] = next
	this.Data["preview"] = preview
	this.TplNames = "artview.html"
}
