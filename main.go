package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/juanfgs/blog/models"
	_ "github.com/juanfgs/blog/routers"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
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
	beego.AddFuncMap("renderPost", renderPost)
	beego.AddFuncMap("renderSafeMarkDown", renderSafeMarkDown)
	beego.AddFuncMap("renderMarkDown", renderMarkDown)			
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

func renderMarkDown(input string) string {
	
	unsafe := blackfriday.MarkdownCommon([]byte(input))
//	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(unsafe)
}
func renderSafeMarkDown(input string) string {
	
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html)
}
func renderPost(content string, contentType string) string {
	if contentType == "markdown" {
		return renderMarkDown(content)
	} else {
		return content
	}
}

