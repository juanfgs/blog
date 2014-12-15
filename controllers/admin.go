package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/utils/pagination"
	"time"
	"log"
)

type AdminController struct {
	MainController
}

func (this *AdminController) Index() {
	this.Layout = "admin/index.tpl"

	var posts []models.Post
	o := orm.NewOrm()
	postsPerPage := 10
	countPosts, err := o.QueryTable("posts").Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)
	o.QueryTable("posts").Limit(postsPerPage, paginator.Offset()).All(&posts)


	this.Data["posts"] = posts

	this.TplNames = "admin/dashboard.tpl"
}

func (this *AdminController) CategoryIndex() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Category List"
	var categories []models.Category
	o := orm.NewOrm()
	categoriesPerPage := 10
	countCategories, err := o.QueryTable("categories").Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, categoriesPerPage, countCategories)
	o.QueryTable("categories").Limit(categoriesPerPage, paginator.Offset()).All(&categories)



	this.Data["Categories"] = categories

	this.TplNames = "admin/categories.tpl"
}

func (this *AdminController) NewCategory(){
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new category"
	this.TplNames = "admin/newcategory.tpl"
}

func (this *AdminController) EditCategory(){
	this.Layout = "admin/index.tpl"
	categoryId, err:= this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	category := new(models.Category)

	o.QueryTable("categories").Filter("id", categoryId).One(category)
	this.Data["Title"] = "Editing Category '"+ category.Title +"'"
	this.Data["Category"] = category
	this.TplNames = "admin/editcategory.tpl"
}

func (this *AdminController) NewCategoryWrite(){
	flash := beego.NewFlash()
	o := orm.NewOrm()

	category := new(models.Category)
	category.Title = this.GetString("Title")
	category.Description = this.GetString("Description")




	_, err := o.Insert(category)
	if err != nil {
		log.Println(err)
		flash.Error("Error creating category")
		flash.Store(&this.Controller)
		this.Abort("500")
		return
	}
	flash.Notice("Category Created")
	flash.Store(&this.Controller)
	this.Redirect("/admin/categories/", 302)
}

func (this *AdminController) EditCategoryWrite(){
	categoryId, err:= this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}

	category := new(models.Category)

	o := orm.NewOrm()


	o.QueryTable("categories").Filter("id", categoryId).One(category)
	if val := this.GetString("Title"); val != category.Title {
		category.Title = val
	}

	if val := this.GetString("Description"); val != category.Description {
		category.Description = val
	}

	if _, err := o.Update(category); err == nil {
		flash.Notice("Category Saved")
		flash.Store(&this.Controller)
		this.Redirect("/admin/categories/", 302)
		return
	} else {

		this.Abort("500")
	}

}
func (this *AdminController) DeletePost(){

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

func (this *AdminController) DeleteCategory(){

	categoryId, err:= this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	category := new(models.Category)
	o.QueryTable("posts").Filter("id", categoryId).One(category)
	if _, err = o.Delete(category); err == nil {
		flash.Notice("Category erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete category, check log for details")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/", 302)
	return
}

func (this *AdminController) EditPost(){
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

func (this *AdminController) EditPostWrite(){
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



func (this *AdminController) NewPost(){
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new post"
	o := orm.NewOrm()
	var categories []models.Category 
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = categories
	this.TplNames = "admin/newpost.tpl"
}

func (this *AdminController) NewPostWrite(){
	flash := beego.NewFlash()
	o := orm.NewOrm()

	post := new(models.Post)
	post.Title = this.GetString("Title")
	post.Tagline = this.GetString("Tagline")
	post.Content = this.GetString("Content")

	post.Description = this.GetString("Keywords")
	post.Description = this.GetString("Keywords")
	published, errbool := this.GetBool("Published")
	if errbool == nil {
		post.Published = published
	} else {
		post.Published = false
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

func (this *AdminController) URLMapping() {
	this.Mapping("Index", this.Index)
	this.Mapping("NewPost", this.NewPost)
	this.Mapping("NewPostWrite", this.NewPostWrite)

}





