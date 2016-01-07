package controllers

import (

	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"

	"log"
)

type PagesController struct {
	MainController
}


// @router /page/:id [get]
func (this *PagesController) Show() {
	this.Layout = "index.tpl"

	id, err := this.GetInt(":id")
	slug := this.GetString(":slug")
	log.Println(slug)
	if err != nil && slug == "" {
		log.Println(id)
		log.Println(err)
		this.Abort("400")
	}

	var  page models.Page

	o := orm.NewOrm()

	if slug == "" {
		err = o.QueryTable("pages").Filter("id", id).One(&page)
	} else {
		err = o.QueryTable("pages").Filter("slug", slug).One(&page)
	}

	this.Data["Name"] = page.Title

	this.Data["Page"] = page


	this.TplNames = "pages/view.tpl"
}

func (this *PagesController) URLMapping() {
	this.Mapping("Show", this.Show)


}
