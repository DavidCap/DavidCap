package controllers

import (
	"DavidCap/models"

	"strings"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	this.Layout = "layout.html"
	this.Data["IsLogin"] = IsLogin(this.Ctx)

	blogs, err := models.GetAllBlog()
	for i, value := range blogs {
		//去除正文中的#
		lenth := len(value.Content)
		if lenth > 50 {
			lenth = 50
		}
		blogs[i].Content = strings.Replace(value.Content[:lenth], "#", "", lenth)
	}

	if err == nil {
		this.Data["Blogs"] = blogs
	}
}
