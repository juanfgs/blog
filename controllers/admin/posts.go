package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/juanfgs/blog/controllers"	
	"github.com/astaxie/beego/utils/pagination"
	"time"
	"log"
)

type PostsController struct {
	controllers.MainController
}

func (this *PostsController) Index() {
	this.Layout = "admin/index.tpl"

	var posts []models.Post
	o := orm.NewOrm()
	postsPerPage := 10
	countPosts, err := o.QueryTable("posts").Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)
	o.QueryTable("posts").Limit(postsPerPage, paginator.Offset()).OrderBy("-created_at").All(&posts)


	this.Data["posts"] = posts

	this.TplNames = "admin/dashboard.tpl"
}

func (this *PostsController) Delete(){

	postId, err:= this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	post := new(models.Post)
	o.QueryTable("posts").Filter("id", postId).One(post)
	if _, err = o.Delete(post); err == nil {
		flash.Notice("Post erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete post")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/", 302)
	return

}

func (this *PostsController) Edit(){
	this.Layout = "admin/index.tpl"
	postId, err:= this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	post := new(models.Post)

	var categories []models.Category 
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = categories

	o.QueryTable("posts").Filter("id", postId).One(post)
	this.Data["Title"] = "Editing Post '"+ post.Title +"'"
	this.Data["Post"] = post

	
	this.TplNames = "admin/editpost.tpl"
}

func (this *PostsController) EditWrite(){
	postId, err:= this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}

	post := new(models.Post)

	o := orm.NewOrm()


	o.QueryTable("posts").Filter("id", postId).One(post)
	if val := this.GetString("Title"); val != post.Title {
		post.Title = val
	}
	if val := this.GetString("Tagline"); val != post.Tagline {
		post.Tagline = val
	}
	if val := this.GetString("Content"); val != post.Content {
		post.Content = val
	}

	if val := this.GetString("Keywords"); val != post.Keywords {
		post.Keywords = val
	}
	if val := this.GetString("ContentType"); val != post.ContentType {
		post.ContentType = val
	}
	
	if val := this.GetString("Description"); val != post.Description {
		post.Description = val
	}

	if val, err := this.GetInt("CategoryId"); err == nil {
			var category models.Category
			o.QueryTable("categories").Filter("id", val).One(&category)
			post.Category = &category
	}
	
	published, errbool := this.GetBool("Published")
	if errbool == nil {
		post.Published = published
	} else {
		post.Published = false
	}
	if _, err := o.Update(post); err == nil {
		flash.Notice("Post Saved")
		flash.Store(&this.Controller)
		this.Redirect("/admin/", 302)
		return
	} else {

		this.Abort("500")
	}

}


func (this *PostsController) New(){
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new post"
	o := orm.NewOrm()
	var categories []models.Category 
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = categories
	this.TplNames = "admin/newpost.tpl"
}

func (this *PostsController) NewWrite(){
	flash := beego.NewFlash()
	o := orm.NewOrm()

	post := new(models.Post)
	post.Title = this.GetString("Title")
	post.Tagline = this.GetString("Tagline")
	post.Content = this.GetString("Content")

	post.Description = this.GetString("Description")
	post.Keywords = this.GetString("Keywords")
	published, errbool := this.GetBool("Published")
	if errbool == nil {
		post.Published = published
	} else {
		post.Published = false
	}
	if val := this.GetString("ContentType"); val != post.ContentType {
		post.ContentType = val
	}

	if val, err := this.GetInt("CategoryId"); err == nil {
			var category models.Category
			o.QueryTable("categories").Filter("id", val).One(&category)
			post.Category = &category
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	if user, ok := this.Data["User"].(models.User); ok {
		post.Author = &user
	}


	_, err := o.Insert(post)
	if err != nil {
		log.Println(err)
		flash.Error("Error creating post")
		flash.Store(&this.Controller)
		this.Abort("500")
		return
	}
	flash.Notice("Post Created")
	flash.Store(&this.Controller)
	this.Redirect("/admin/", 302)
}
