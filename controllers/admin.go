package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"time"
)

type AdminController struct {
	MainController
}

func (this *AdminController) Index() {
	this.Layout = "admin/index.tpl"

	var posts []models.Post
	o := orm.NewOrm()

	o.QueryTable("posts").All(&posts)

	this.Data["posts"] = posts

	this.TplNames = "admin/dashboard.tpl"
}

func (this *AdminController) NewPost(){
	this.Layout = "admin/index.tpl"

	this.TplNames = "admin/newpost.tpl"
}

func (this *AdminController) NewPostWrite(){
	this.Layout = "admin/index.tpl"
	o := orm.NewOrm()

	post := new(models.Post)
	post.Title = this.GetString("Title")
	post.Tagline = this.GetString("Tagline")
	post.Content = this.GetString("Content")
	post.Published = true
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	o.Insert(post)

	this.Redirect("/admin/", 302)
}

func (this *AdminController) URLMapping() {
	this.Mapping("Index", this.Index)
	this.Mapping("NewPost", this.NewPost)
	this.Mapping("NewPostWrite", this.NewPostWrite)

}




