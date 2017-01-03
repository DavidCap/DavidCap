package routers

import (
	"DavidCap/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/addBlog", &controllers.AddBlogController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.AutoRouter(&controllers.BlogController{})
	beego.Router("/aboutMe", &controllers.AboutMeController{})
	beego.Router("/login", &controllers.LoginController{})
}
