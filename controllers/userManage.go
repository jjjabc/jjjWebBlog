package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jjjabc/jjjWebBlog/user"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	juId, err := strconv.Atoi(this.GetString("userId"))
	if err != nil {
		this.Ctx.WriteString("Get juId error!")
		return
	}
	ju := user.GetUser(juId)
	if ju != nil {
		this.Ctx.WriteString("Get ju-object error!")
		return
	}
	this.Data["user"] = ju
	this.TplNames = "userView.html"
}
func (this *UserController) Post() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	ju := user.JJJuser{}
	ju.Name = this.GetString("name")
	ju.NickName = this.GetString("nickname")
	ju.Description = this.GetString("description")
	if this.GetString("password") == "" {
		this.Ctx.WriteString("Password is empty")
		return
	}

	if err := ju.SigupUser(this.GetString("password")); err != nil {
		this.Ctx.WriteString("Sigupuser error")
		return
	}
	this.Ctx.WriteString("OK")
}
func (this *UserController) Delete() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	juId, err := strconv.Atoi(this.GetString("userId"))
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	ju := user.GetUser(juId)
	if ju == nil {
		this.Ctx.WriteString("Get ju-object error" + strconv.Itoa(juId))
		return
	}
	if err := ju.DelUser(); err != nil {
		this.Ctx.WriteString("Delete user error:" + err.Error())
		return
	}
	this.Ctx.WriteString("OK")
	return
}
func (this *UserController) Put() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	juId, err := strconv.Atoi(this.GetString("userid"))
	if err != nil {
		this.Ctx.WriteString("Get juId=" + this.GetString("userId") + " error:" + err.Error())
		return
	}
	ju := user.GetUser(juId)
	if ju == nil {
		this.Ctx.WriteString("Get user error")
		return
	}
	ju.Name = this.GetString("name")
	ju.Description = this.GetString("des")
	if err := ju.Updata(); err != nil {
		beego.Info("Updata error")
		this.Ctx.WriteString("updata error")
		return
	}
	this.Ctx.WriteString("OK")
}
func (this *UserController) List() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.WriteString("Not login！")
		return
	}
	jus, err := user.GetAllUsers()
	if err != nil {
		beego.Info("error!")
		this.Ctx.WriteString("Get articles error!")
	}
	beego.Info("jus len:" + strconv.Itoa(len(jus)))
	this.Data["jus"] = jus
	this.TplNames = "userList.html"
}
