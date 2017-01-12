package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/models"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare(){
	_ = beego.ReadFromRequest(&this.Controller)
	var sessionName = beego.AppConfig.String("SessionName")
	v := this.GetSession(sessionName)

	o := orm.NewOrm()
	if v != nil { //user logged in
		user, err:= models.NewUser(v.(int))
		if err != nil {
			this.Abort("401")
			return
		}
		this.Data["User"] = user

		// Load notifications
		
		var notifications []models.Notification
		o.QueryTable("notifications").Filter("status", models.NotificationStatusNew).OrderBy("-created_at").Limit(5).All(&notifications)
		cnt, err := o.QueryTable("notifications").Filter("status", models.NotificationStatusNew).Count()
		
		if err != nil {
			panic(err)
		}

		this.Data["HasNotifications"] = (len(notifications) > 0)
		this.Data["NotificationCount"] = cnt
		this.Data["Notifications"] = &notifications
	}
	var categories []models.Category
	o.QueryTable("categories").All(&categories)

	var pages []models.Page
	o.QueryTable("pages").All(&pages)
	var links []models.Link
	o.QueryTable("links").All(&links)

	this.Data["Categories"] = &categories
	this.Data["Pages"] = &pages

	this.Data["Links"] = &links


}



