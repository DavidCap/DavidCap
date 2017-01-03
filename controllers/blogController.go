package controllers

import (
	"DavidCap/models"
	"html"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Index() {
	this.Layout = "layout.html"
	this.TplName = "blog.html"
	this.Data["IsLogin"] = IsLogin(this.Ctx)

	id := this.Ctx.Input.Param("0")
	blog, err := models.GetBlogById(id)
	if err != nil {
		return
	}
	unsafe := blackfriday.MarkdownCommon([]byte(blog.Content))
	contentMDHTML := html.UnescapeString(string(unsafe))
	blog.Content = contentMDHTML

	this.Data["Blog"] = blog
}
