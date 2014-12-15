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
	
	if v != nil { //user logged in
		user, err:= models.NewUser(v.(int))
		if err != nil {
			this.Abort("401")
			return
		}
		this.Data["User"] = user
	}
	o := orm.NewOrm()
	var categories []models.Category
	o.QueryTable("categories").All(&categories)
	this.Data["Categories"] = &categories



}



