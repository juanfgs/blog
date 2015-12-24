package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/juanfgs/blog/controllers"
)

type SettingsController struct {
	controllers.MainController
}


func (this *SettingsController) Index() {
	this.Layout = "admin/index.tpl"
	var settings []models.Setting
	o := orm.NewOrm()
	o.QueryTable("settings").All(&settings)

	this.Data["settings"] = settings
	this.TplNames = "admin/settings.tpl"
}

func (this *SettingsController) Save(){
	flash := beego.NewFlash()
	flash.Notice("Settings saved")
	flash.Store(&this.Controller)
}
