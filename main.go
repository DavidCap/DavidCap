package main

import (
	_ "DavidCap/models"
	_ "DavidCap/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
