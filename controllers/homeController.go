package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "index.html"
	this.Layout = "layout.html"

	this.Data["IsLogin"] = IsLogin(this.Ctx)
}
