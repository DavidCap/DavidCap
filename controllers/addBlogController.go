package controllers

import "github.com/astaxie/beego"
import "DavidCap/models"

type AddBlogController struct {
	beego.Controller
}

func (this *AddBlogController) Get() {
	this.TplName = "addBlog.html"
	this.Layout = "layout.html"
	this.Data["IsLogin"] = IsLogin(this.Ctx)
}

func (this *AddBlogController) Post() {
	if IsLogin(this.Ctx) == false {
		this.Ctx.Redirect(302, "/category")
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	err := models.AddBlog(title, content)
	if err != nil {
		panic(err)
	}
	this.Ctx.Redirect(302, "/category")
}
