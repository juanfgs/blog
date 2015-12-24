package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.PostsController{}, "get:Index")
	beego.Router("/search/", &controllers.PostsController{}, "get:Search")
	beego.Router("/post/:id([0-9]+)", &controllers.PostsController{}, "get:Show")
	beego.Router("/post/:slug([a-z-]+)", &controllers.PostsController{}, "get:Show")

	beego.Router("/comment/new", &controllers.CommentsController{}, "post:CommentWrite")
	beego.Router("/categories/:id([0-9]+)", &controllers.CategoriesController{}, "get:Show")

	beego.Router("/login", &controllers.UsersController{}, "get:Login")
	beego.Router("/login/process", &controllers.UsersController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
	beego.Router("/register/process", &controllers.UsersController{}, "post:RegisterPost")

	beego.Router("/feed.xml", &controllers.FeedsController{}, "get:Index")

	//	beego.Router("/images/:height([0-9])x:width[0-9]/:image:string", &controllers.ThumbnailsController{}, "get:showImage")

}
