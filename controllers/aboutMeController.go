package controllers

import (
	"github.com/astaxie/beego"
)

type AboutMeController struct {
	beego.Controller
}

func (this *AboutMeController) Get() {
	this.Layout = "layout.html"
	this.TplName = "aboutMe.html"
	this.Data["IsLogin"] = IsLogin(this.Ctx)
}
