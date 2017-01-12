package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"
	"log"
)

type LinksController struct {
	controllers.MainController
}

func (this *LinksController) Index() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Link List"
	var links []models.Link
	o := orm.NewOrm()
	linksPerPage := 10
	countLinks, err := o.QueryTable("links").Count()
	if err != nil {
		log.Println(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, linksPerPage, countLinks)
	o.QueryTable("links").Limit(linksPerPage, paginator.Offset()).All(&links)

	this.Data["Links"] = links

	this.TplName = "admin/links/index.tpl"
}

func (this *LinksController) New() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Create new link"
	this.TplName = "admin/links/new.tpl"
}

func (this *LinksController) Edit() {
	this.Layout = "admin/index.tpl"
	linkId, err := this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	link := new(models.Link)

	o.QueryTable("links").Filter("id", linkId).One(link)
	this.Data["Title"] = "Editing Link '" + link.Title + "'"
	this.Data["Link"] = link
	this.TplName = "admin/links/edit.tpl"
}

func (this *LinksController) NewWrite() {
	flash := beego.NewFlash()
	o := orm.NewOrm()

	link := new(models.Link)
	link.Title = this.GetString("Title")
	link.Url = this.GetString("Url")

	_, err := o.Insert(link)
	if err != nil {
		log.Println(err)
		flash.Error("Error creating link")
		flash.Store(&this.Controller)
		this.Abort("500")
		return
	}
	flash.Notice("Link Created")
	flash.Store(&this.Controller)
	this.Redirect("/admin/links/", 302)
}

func (this *LinksController) EditWrite() {
	linkId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}

	link := new(models.Link)

	o := orm.NewOrm()

	o.QueryTable("links").Filter("id", linkId).One(link)
	if val := this.GetString("Title"); val != link.Title {
		link.Title = val
	}

	if val := this.GetString("Url"); val != link.Url {
		link.Url = val
	}

	if _, err := o.Update(link); err == nil {
		flash.Notice("Link Saved")
		flash.Store(&this.Controller)
		this.Redirect("/admin/links/", 302)
		return
	} else {

		this.Abort("500")
	}

}

func (this *LinksController) Delete() {

	linkId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	link := new(models.Link)
	o.QueryTable("links").Filter("id", linkId).One(link)
	if _, err = o.Delete(link); err == nil {
		flash.Notice("Link erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete link, check log for details")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/links", 302)
	return
}
