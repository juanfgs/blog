package controllers


import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego"
	"log"
)

type CategoriesController struct {
	MainController
}


func (this *CategoriesController) Show() {
	this.Layout = "index.tpl"

	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(id)
		log.Println(err)
		this.Abort("400")
	} 

	var category models.Category
	
	o := orm.NewOrm()

	err = o.QueryTable("categories").Filter("id", id).One(&category)

	if err != nil {
		this.Abort("404")
	}

	this.Data["Title"] = "Posts archived in  category: " + category.Title
	this.Data["HeroTitle"] = category.Title
	this.Data["HeroTagline"] = category.Description
	this.Data["Category"] = category


	countPosts, err := o.QueryTable("posts").Filter("category_id", category.Id).Filter("published",1).Count();


	postsPerPage, err := beego.AppConfig.Int("postsperpage") 

	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, countPosts)


	_, err = o.LoadRelated(&category, "Posts",0, postsPerPage , paginator.Offset(), "-created_at")
	if err == nil {
		this.Data["posts"] = category.Posts
	}
	this.TplName = "categories/view.tpl"
}

func (this *CategoriesController) URLMapping() {
	this.Mapping("Show", this.Show)
}
