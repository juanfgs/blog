package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.PostsController{}, "get:Index")
	beego.Router("/post/:id([0-9]+)", &controllers.PostsController{}, "get:Show")
	beego.Router("/comment/new",&controllers.CommentsController{}, "post:CommentWrite")	
	beego.Router("/categories/:id([0-9]+)", &controllers.CategoriesController{}, "get:Show")
	beego.Router("/admin/", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/post/new", &controllers.AdminController{}, "get:NewPost")
	beego.Router("/admin/post/new", &controllers.AdminController{}, "post:NewPostWrite")
	beego.Router("/admin/post/edit/:id([0-9]+)", &controllers.AdminController{}, "get:EditPost")
	beego.Router("/admin/post/edit/:id([0-9]+)", &controllers.AdminController{}, "post:EditPostWrite")
	beego.Router("/admin/post/delete/:id([0-9]+)", &controllers.AdminController{}, "get:DeletePost")
	beego.Router("/admin/categories/", &controllers.AdminController{}, "get:CategoryIndex")
	beego.Router("/admin/categories/new", &controllers.AdminController{}, "get:NewCategory")
	beego.Router("/admin/categories/new", &controllers.AdminController{}, "post:NewCategoryWrite")
	beego.Router("/admin/categories/edit/:id", &controllers.AdminController{}, "get:EditCategory")
	beego.Router("/admin/categories/edit/:id", &controllers.AdminController{}, "post:EditCategoryWrite")
	beego.Router("/admin/categories/delete/:id([0-9]+)", &controllers.AdminController{}, "get:DeleteCategory")
	beego.Router("/login", &controllers.UsersController{}, "get:Login")
	beego.Router("/login/process", &controllers.UsersController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
	beego.Router("/register/process", &controllers.UsersController{}, "post:RegisterPost")

//	beego.Router("/images/:height([0-9])x:width[0-9]/:image:string", &controllers.ThumbnailsController{}, "get:showImage")
	
}

