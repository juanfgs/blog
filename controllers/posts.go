package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/utils/pagination"
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
	postsPerPage, err := beego.AppConfig.Int("postsperpage")
	countPosts, err := o.QueryTable("posts").Filter("published", 1).Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)
	
	o.QueryTable("posts").Filter("published", 1).Limit(postsPerPage, paginator.Offset()).OrderBy("-created_at").All(&posts)

	this.Data["posts"] = posts

	this.TplNames = "posts.tpl"

}


// @router /post/:id [get]
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
	_, err = o.LoadRelated(&post, "Author")

	this.Data["Title"] = post.Title


	this.Data["HeroTagline"] = "Mostly programming stuff, but also my life"
	this.Data["Post"] = post
	this.TplNames = "post.tpl"
}

func (this *PostsController) URLMapping() {
	this.Mapping("Show", this.Show)
	this.Mapping("Index", this.Index)

}
