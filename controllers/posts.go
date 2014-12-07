package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"log"
)

type PostsController struct {
	MainController
}

func (this *PostsController) Index() {
	this.Layout = "index.tpl"
	this.Data["Title"] = "Juan Giménez Silva's Blog"
	this.Data["HeroTitle"] = "Welcome"
	this.Data["HeroTagline"] = "Mostly programming stuff, but also my life"
	var posts []models.Post
	o := orm.NewOrm()

	o.QueryTable("posts").Filter("published", 1).All(&posts)

	this.Data["posts"] = posts

	this.TplNames = "posts.tpl"

}


// @router /posts/:id [get]
func (this *PostsController) Show() {
	this.Layout = "index.tpl"

	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(id)
		log.Println(err)
		this.Abort("400")
	} 

	var post models.Post
	
	o := orm.NewOrm()
	err = o.QueryTable("posts").Filter("id", id).One(&post)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Post"] = post
	this.TplNames = "post.tpl"
}

func (this *PostsController) URLMapping() {
	this.Mapping("Show", this.Show)
	this.Mapping("Index", this.Index)

}
