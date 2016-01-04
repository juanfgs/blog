package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers/admin"
)

func init() {
	beego.Router("/admin/", &admin.PostsController{}, "get:Index")
	beego.Router("/admin/posts", &admin.PostsController{}, "get:Index")
	beego.Router("/admin/posts/new", &admin.PostsController{}, "get:New")
	beego.Router("/admin/posts/new", &admin.PostsController{}, "post:NewWrite")
	beego.Router("/admin/posts/edit/:id([0-9]+)", &admin.PostsController{}, "get:Edit")
	beego.Router("/admin/posts/edit/:id([0-9]+)", &admin.PostsController{}, "post:EditWrite")
	beego.Router("/admin/posts/delete/:id([0-9]+)", &admin.PostsController{}, "get:Delete")

	// Pages
	beego.Router("/admin/pages", &admin.PagesController{}, "get:Index")
	beego.Router("/admin/pages/new", &admin.PagesController{}, "get:New")
	beego.Router("/admin/pages/new", &admin.PagesController{}, "post:NewWrite")
	beego.Router("/admin/pages/edit/:id([0-9]+)", &admin.PagesController{}, "get:Edit")
	beego.Router("/admin/pages/edit/:id([0-9]+)", &admin.PagesController{}, "post:EditWrite")
	beego.Router("/admin/pages/delete/:id([0-9]+)", &admin.PagesController{}, "get:Delete")

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
	beego.Router("/admin/comments/edit/:id([0-9]+)", &admin.CommentsController{}, "get:Edit")
	beego.Router("/admin/comments/edit/:id([0-9]+)", &admin.CommentsController{}, "post:EditWrite")

	// Settings route
	beego.Router("/admin/settings/", &admin.SettingsController{}, "get:Index")
	beego.Router("/admin/settings/save", &admin.SettingsController{}, "post:Save")	
}
