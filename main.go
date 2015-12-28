package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/juanfgs/blog/helpers"
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

	beego.AddFuncMap("equals", equals)
	beego.AddFuncMap("renderPost", helpers.RenderPost)
	beego.AddFuncMap("renderSafeMarkDown", helpers.RenderSafeMarkDown)
	beego.AddFuncMap("renderMarkDown", helpers.RenderMarkDown)
	beego.InsertFilter("/admin/", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

func equals(a interface{}, b interface{}) bool {
	if a == nil || b == nil {
		return false
	}

	if a == b {
		return true
	}
	return false
}
