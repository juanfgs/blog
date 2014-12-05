package controllers

import (
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/orm"
)

type PostsController struct {
	MainController
}

func (this *PostsController) Get(){
	this.Data["Title"] = "Juan Giménez Silva's Blog"
	this.Data["HeroTitle"] = "Welcome"
	this.Data["HeroTagline"] = "Mostly programming stuff, but also my life"

	var posts []models.Post
	o := orm.NewOrm()



	o.QueryTable("posts").Filter("published", 1).All(&posts)

	this.Data["posts"] = posts

	this.TplNames = "index.tpl"

}


