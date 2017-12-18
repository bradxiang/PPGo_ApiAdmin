package main

import (
	"github.com/api-admin/models"
	_ "github.com/api-admin/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
