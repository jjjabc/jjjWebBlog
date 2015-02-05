package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"os"
	"strconv"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

func (this *UploadFileController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	remoteFile, fileHeader, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		this.Ctx.WriteString("error")
		return
	}
	defer remoteFile.Close()
	path := "./static/img/upload/"
	remoteFilename := fileHeader.Filename
	filename := strconv.FormatInt(time.Now().Unix(), 10) + remoteFilename
	f, err := os.Create(path + filename)
	if err != nil {
		beego.Info("create file err:" + err.Error())
	}
	defer f.Close()
	io.Copy(f, remoteFile)
	this.Ctx.WriteString(path + filename)
	return
}
