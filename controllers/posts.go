package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/utils/pagination"
	"log"
	"fmt"
)

type PostsController struct {
	MainController
}

func (this *PostsController) Index() {
	this.Layout = "index.tpl"
	this.Data["Title"] = "Juan Giménez Silva's Blog"
	this.Data["HeroTitle"] = "Bienvenidos a mi blog"
	this.Data["HeroTagline"] = "Compartiendo un poco de lo que no sé"
	var posts []models.Post
	o := orm.NewOrm()
	postsPerPage, err := beego.AppConfig.Int("postsperpage")
	countPosts, err := o.QueryTable("posts").Filter("published", 1).Count()

	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)

	o.QueryTable("posts").Filter("published", 1).Limit(postsPerPage, paginator.Offset()).OrderBy("-created_at").All(&posts)
	for idx,_ := range posts {
		_, err = o.LoadRelated(&posts[idx], "Comments")
	}
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

	_, err = o.LoadRelated(&post, "Comments")
	if err == nil {
		this.Data["Comments"] = post.Comments
	} else {
		log.Println(err)
	}

	this.TplNames = "post.tpl"
}

func (this *PostsController) Search(){
	this.Layout = "index.tpl"
	keyword := this.GetString("keyword")
	var posts []models.Post

	o := orm.NewOrm()

	postsPerPage, err := beego.AppConfig.Int("postsperpage")

	countPosts, err := o.QueryTable("posts").Filter("title__icontains", keyword).Count()

	if err != nil {
		log.Println(err)
	}

	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)
	this.Data["Title"] = "Search:" + keyword
	this.Data["HeroTitle"] = "Search:" + keyword

	keyword = "%"+keyword+"%"


	num, err := o.Raw("SELECT * FROM `posts` WHERE published = 1 AND (title LIKE ? OR content LIKE ?) ORDER BY created_at DESC LIMIT ? OFFSET ? ", keyword,keyword,postsPerPage,paginator.Offset()).QueryRows(&posts)


	this.Data["HeroTagline"] = fmt.Sprintf( "Entries: %d", num)
	if err != nil {
		log.Println(err)
	}
	for idx,_ := range posts {
		_, err = o.LoadRelated(&posts[idx], "Comments")
	}

	this.Data["posts"] = posts

	this.TplNames = "posts.tpl"
}

func (this *PostsController) URLMapping() {
	this.Mapping("Show", this.Show)
	this.Mapping("Index", this.Index)
	this.Mapping("Search", this.Search)
}
