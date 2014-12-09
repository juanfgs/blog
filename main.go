package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/juanfgs/blog/models"
	_ "github.com/juanfgs/blog/routers"
)
var sessionName = beego.AppConfig.String("SessionName")
func main() {

	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session(sessionName).(int)
		if !ok && ctx.Input.Uri() != "/login" && ctx.Input.Uri() != "/register" {
			ctx.Redirect(302, "/login")
		}
	}
	

	
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}
