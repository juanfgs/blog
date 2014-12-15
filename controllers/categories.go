package controllers


import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"

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
	_, err = o.LoadRelated(&category,"Posts")
	if err == nil {
		this.Data["posts"] = category.Posts
	}
	this.TplNames = "category.tpl"
}

func (this *CategoriesController) URLMapping() {
	this.Mapping("Show", this.Show)


}
