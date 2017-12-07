package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"
	"log"
)

type CategoriesController struct {
	controllers.MainController
}

func (this *CategoriesController) Index() {
	var categories []models.Category
	o := orm.NewOrm()
	_, err := o.QueryTable("categories").Count()
	if err != nil {
		log.Println(err)
	}
	o.QueryTable("categories").All(&categories)

	this.Data["json"] = categories

	this.ServeJSON()
}

func (this *CategoriesController) New() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new category"
	this.TplName = "admin/categories/new.tpl"
}

func (this *CategoriesController) Edit() {
	this.Layout = "admin/index.tpl"
	categoryId, err := this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	category := new(models.Category)

	o.QueryTable("categories").Filter("id", categoryId).One(category)
	this.Data["Title"] = "Editing Category '" + category.Title + "'"
	this.Data["Category"] = category
	this.TplName = "admin/categories/edit.tpl"
}

func (this *CategoriesController) NewWrite() {
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

func (this *CategoriesController) EditWrite() {
	categoryId, err := this.GetInt(":id")
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

func (this *CategoriesController) Delete() {

	categoryId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	category := new(models.Category)
	o.QueryTable("categories").Filter("id", categoryId).One(category)
	if _, err = o.Delete(category); err == nil {
		flash.Notice("Category erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete category, check log for details")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/categories", 302)
	return
}
