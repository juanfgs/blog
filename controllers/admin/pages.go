package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"
	"log"
	"time"
)

type PagesController struct {
	controllers.MainController
}

func (this *PagesController) Index() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Pages"
	var pages []models.Page
	o := orm.NewOrm()
	pagesPerPage := 10
	countPages, err := o.QueryTable("pages").Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, pagesPerPage, countPages)
	o.QueryTable("pages").Limit(pagesPerPage, paginator.Offset()).OrderBy("-created_at").All(&pages)

	this.Data["pages"] = pages

	this.TplNames = "admin/pages/index.tpl"
}

func (this *PagesController) Delete() {

	pageId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	page := new(models.Page)
	o.QueryTable("pages").Filter("id", pageId).One(page)
	if _, err = o.Delete(page); err == nil {
		flash.Notice("Page erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete page")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/pages/", 302)
	return

}

func (this *PagesController) Edit() {
	this.Layout = "admin/index.tpl"
	pageId, err := this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	page := new(models.Page)

	var categories []models.Category
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = categories

	o.QueryTable("pages").Filter("id", pageId).One(page)
	this.Data["Title"] = "Editing Page '" + page.Title + "'"
	this.Data["Page"] = page

	this.TplNames = "admin/pages/edit.tpl"
}

func (this *PagesController) EditWrite() {
	pageId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}

	page := new(models.Page)

	o := orm.NewOrm()

	o.QueryTable("pages").Filter("id", pageId).One(page)
	if val := this.GetString("Title"); val != page.Title {
		page.Title = val
	}
	if val := this.GetString("Tagline"); val != page.Tagline {
		page.Tagline = val
	}
	if val := this.GetString("Content"); val != page.Content {
		page.Content = val
	}

	if val := this.GetString("Keywords"); val != page.Keywords {
		page.Keywords = val
	}
	if val := this.GetString("ContentType"); val != page.ContentType {
		page.ContentType = val
	}

	if val := this.GetString("Description"); val != page.Description {
		page.Description = val
	}

	if val := this.GetString("Slug"); val == "" {
		page.Slug = string(models.GenerateSlug(page.Title))
	}



	published, errbool := this.GetBool("Published")
	if errbool == nil {
		page.Published = published
	} else {
		page.Published = false
	}
	if _, err := o.Update(page); err == nil {
		flash.Notice("Page Saved")
		flash.Store(&this.Controller)
		this.Redirect("/admin/pages/", 302)
		return
	} else {

		this.Abort("500")
	}

}

func (this *PagesController) New() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new page"
	o := orm.NewOrm()
	var categories []models.Category
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = categories
	this.TplNames = "admin/pages/new.tpl"
}

func (this *PagesController) NewWrite() {
	flash := beego.NewFlash()
	o := orm.NewOrm()

	page := new(models.Page)
	page.Title = this.GetString("Title")
	page.Tagline = this.GetString("Tagline")
	page.Content = this.GetString("Content")

	page.Description = this.GetString("Description")
	page.Keywords = this.GetString("Keywords")
	published, errbool := this.GetBool("Published")
	if errbool == nil {
		page.Published = published
	} else {
		page.Published = false
	}
	if val := this.GetString("ContentType"); val != page.ContentType {
		page.ContentType = val
	}

	if val := this.GetString("Slug"); val == "" {
		page.Slug = string(models.GenerateSlug(page.Title))
	}
	
	page.CreatedAt = time.Now()
	page.UpdatedAt = time.Now()


	_, err := o.Insert(page)
	if err != nil {
		log.Println(err)
		flash.Error("Error creating page")
		flash.Store(&this.Controller)
		this.Abort("500")
		return
	}
	flash.Notice("Page Created")
	flash.Store(&this.Controller)
	this.Redirect("/admin/pages/", 302)
}
