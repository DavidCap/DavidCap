package controllers

import (
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	if IsLogin(this.Ctx) {
		this.Ctx.Redirect(302, "/")
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	userName := this.Input().Get("uname")
	password := this.Input().Get("pwd")

	if CheckLogin(userName, password) {
		maxAge := 1<<31 - 1
		this.Ctx.SetCookie("userName", userName, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
		this.Ctx.Redirect(302, "/login")
		return
	}
	this.Ctx.Redirect(302, "/")
}

func IsLogin(ctx *context.Context) bool {
	userName, err := ctx.Request.Cookie("userName")
	if err != nil {
		return false
	}

	password, err := ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	return CheckLogin(userName.Value, password.Value)
}

func CheckLogin(userName, password string) bool {
	return userName == beego.AppConfig.String("userName") && password == beego.AppConfig.String("password")
}
