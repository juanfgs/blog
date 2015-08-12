package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.PostsController{}, "get:Index")
	beego.Router("/search/", &controllers.PostsController{}, "get:Search")
	beego.Router("/post/:id([0-9]+)", &controllers.PostsController{}, "get:Show")

	beego.Router("/comment/new", &controllers.CommentsController{}, "post:CommentWrite")
	beego.Router("/categories/:id([0-9]+)", &controllers.CategoriesController{}, "get:Show")
	beego.Router("/admin/", &admin.PostsController{}, "get:Index")
	beego.Router("/admin/post/new", &admin.PostsController{}, "get:New")
	beego.Router("/admin/post/new", &admin.PostsController{}, "post:NewWrite")
	beego.Router("/admin/post/edit/:id([0-9]+)", &admin.PostsController{}, "get:Edit")
	beego.Router("/admin/post/edit/:id([0-9]+)", &admin.PostsController{}, "post:EditWrite")
	beego.Router("/admin/post/delete/:id([0-9]+)", &admin.PostsController{}, "get:Delete")

	//Media routes
	beego.Router("/admin/media/create", &admin.MediaController{}, "post:Create")
	beego.Router("/admin/media/upload", &admin.MediaController{}, "get:Upload")
	beego.Router("/admin/media/", &admin.MediaController{}, "get:Index")
	beego.Router("/admin/media/delete/:id([0-9]+)", &admin.MediaController{}, "get:Delete")

	beego.Router("/admin/categories/", &admin.CategoriesController{}, "get:Index")
	beego.Router("/admin/categories/new", &admin.CategoriesController{}, "get:New")
	beego.Router("/admin/categories/new", &admin.CategoriesController{}, "post:NewWrite")
	beego.Router("/admin/categories/edit/:id", &admin.CategoriesController{}, "get:Edit")
	beego.Router("/admin/categories/edit/:id", &admin.CategoriesController{}, "post:EditWrite")
	beego.Router("/admin/categories/delete/:id([0-9]+)", &admin.CategoriesController{}, "get:Delete")

	beego.Router("/admin/comments/", &admin.CommentsController{}, "get:Index")
	beego.Router("/admin/comments/delete/:id([0-9]+)", &admin.CommentsController{}, "get:Delete")	


	beego.Router("/login", &controllers.UsersController{}, "get:Login")
	beego.Router("/login/process", &controllers.UsersController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
	beego.Router("/register/process", &controllers.UsersController{}, "post:RegisterPost")

	//	beego.Router("/images/:height([0-9])x:width[0-9]/:image:string", &controllers.ThumbnailsController{}, "get:showImage")

}
