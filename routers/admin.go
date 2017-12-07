package routers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers/admin"
)

func init() {
	//dashboard
	beego.Router("/admin/", &admin.DashboardController{}, "get:Index")
	
	beego.Router("/admin/posts/:id([0-9]+)", &admin.PostsController{}, "get:GetPost")
	beego.Router("/admin/posts/", &admin.PostsController{}, "get:GetPosts")
	beego.Router("/admin/posts/:id([0-9]+)", &admin.PostsController{}, "put:PatchPost")
	beego.Router("/admin/posts/", &admin.PostsController{}, "put:PutPost")
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
	beego.Router("/admin/categories/:id([0-9])", &admin.CategoriesController{}, "get:Edit")
	beego.Router("/admin/categories/edit/:id", &admin.CategoriesController{}, "post:EditWrite")
	beego.Router("/admin/categories/delete/:id([0-9]+)", &admin.CategoriesController{}, "get:Delete")

	beego.Router("/admin/categories/new", &admin.CategoriesController{}, "post:NewWrite")

	beego.Router("/admin/links/", &admin.LinksController{}, "get:Index")
	beego.Router("/admin/links/new", &admin.LinksController{}, "get:New")
	beego.Router("/admin/links/new", &admin.LinksController{}, "post:NewWrite")
	beego.Router("/admin/links/edit/:id", &admin.LinksController{}, "get:Edit")
	beego.Router("/admin/links/edit/:id", &admin.LinksController{}, "post:EditWrite")
	beego.Router("/admin/links/delete/:id([0-9]+)", &admin.LinksController{}, "get:Delete")

	
	beego.Router("/admin/comments/", &admin.CommentsController{}, "get:Index")
	beego.Router("/admin/comments/delete/:id([0-9]+)", &admin.CommentsController{}, "get:Delete")
	beego.Router("/admin/comments/edit/:id([0-9]+)", &admin.CommentsController{}, "get:Edit")
	beego.Router("/admin/comments/edit/:id([0-9]+)", &admin.CommentsController{}, "post:EditWrite")

	// Settings route
	beego.Router("/admin/settings/", &admin.SettingsController{}, "get:Index")
	beego.Router("/admin/settings/save", &admin.SettingsController{}, "post:Save")	
}
