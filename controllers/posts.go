package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/juanfgs/blog/models"
	"log"
)

type PostsController struct {
	MainController
}

func (this *PostsController) Index() {
	this.Layout = "index.tpl"

	if title, err := models.GetSetting("Title"); err == nil {
		this.Data["Title"] = title.Value
	}
	if mainTitle, err := models.GetSetting("MainTitle"); err == nil {
		this.Data["MainTitle"] = mainTitle.Value
	}
	if secondaryTitle, err := models.GetSetting("SecondaryTitle"); err == nil {
		this.Data["SecondaryTitle"] = secondaryTitle.Value
	}

	if prettyUrls, err := models.GetSetting("EnablePrettyUrls"); err == nil {
		this.Data["PrettyUrls"] = prettyUrls.Value
	}

	var posts []models.Post
	o := orm.NewOrm()
	postsPerPage, err := beego.AppConfig.Int("postsperpage")
	countPosts, err := o.QueryTable("posts").Filter("published", 1).Count()

	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)

	o.QueryTable("posts").Filter("published", 1).Limit(postsPerPage, paginator.Offset()).OrderBy("-created_at").All(&posts)
	for idx, _ := range posts {
		_, err = o.LoadRelated(&posts[idx], "Comments")
	}
	this.Data["posts"] = posts

	this.TplNames = "posts/index.tpl"

}

// @router /post/:id [get]
func (this *PostsController) Show() {
	this.Layout = "index.tpl"

	id, err := this.GetInt(":id")
	slug := this.GetString(":slug")

	if err != nil && slug == "" {
		log.Println(id)
		log.Println(err)
		this.Abort("400")
	}

	var post models.Post

	o := orm.NewOrm()

	if slug == "" {
		err = o.QueryTable("posts").Filter("id", id).One(&post)
	} else {
		err = o.QueryTable("posts").Filter("slug", slug).One(&post)
	}
	_, err = o.LoadRelated(&post, "Author")

	this.Data["Title"] = post.Title
	this.Data["SecondaryTitle"] = post.Tagline
	this.Data["Post"] = post
	if enabled, err := models.GetSetting("EnableComments"); err == nil {
		this.Data["EnableComments"] = enabled.Value
	}

	_, err = o.LoadRelated(&post, "Comments")

	if err == nil {
		for idx, _ := range post.Comments {
			_, err = o.LoadRelated(post.Comments[idx], "Parent")
		}

		this.Data["Comments"] = post.Comments
	} else {
		log.Println(err)
	}

	this.TplNames = "posts/view.tpl"
}

func (this *PostsController) Search() {
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
	this.Data["MainTitle"] = "Search:" + keyword

	keyword = "%" + keyword + "%"

	num, err := o.Raw("SELECT * FROM `posts` WHERE published = 1 AND (title LIKE ? OR content LIKE ?) ORDER BY created_at DESC LIMIT ? OFFSET ? ", keyword, keyword, postsPerPage, paginator.Offset()).QueryRows(&posts)

	this.Data["HeroTagline"] = fmt.Sprintf("Entries: %d", num)
	if err != nil {
		log.Println(err)
	}
	for idx, _ := range posts {
		_, err = o.LoadRelated(&posts[idx], "Comments")
	}

	this.Data["posts"] = posts

	this.TplNames = "posts/index.tpl"
}

func (this *PostsController) URLMapping() {
	this.Mapping("Show", this.Show)
	this.Mapping("Index", this.Index)
	this.Mapping("Search", this.Search)
}
