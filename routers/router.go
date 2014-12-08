package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.PostsController{}, "get:Index")
	beego.Router("/post/:id([0-9]+)", &controllers.PostsController{}, "get:Show")
	beego.Router("/admin/", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/post/new", &controllers.AdminController{}, "get:NewPost")
	beego.Router("/admin/post/new", &controllers.AdminController{}, "post:NewPostWrite")


	beego.Router("/login", &controllers.UsersController{}, "get:Login")
	beego.Router("/login/process", &controllers.UsersController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
	beego.Router("/register/process", &controllers.UsersController{}, "post:RegisterPost")


}

