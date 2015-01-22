package controllers

import (
	"github.com/astaxie/beego"
	"jjjBlog/article"
)

type AddartController struct {
	beego.Controller
}

func (c *AddartController) Get() {
	ja := article.JJJarticle{
		Title: "hello",
		Text:  "world!",
	}
	if err := ja.AddArticle(); err != nil {
		beego.Info("add error")
	}
	if err := ja.Publish(); err != nil {
		beego.Info("publish error")
	}
	c.Ctx.Redirect(302, "/list")
}
