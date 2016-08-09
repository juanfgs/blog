package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/juanfgs/blog/controllers"
	"log"
)

type SettingsController struct {
	controllers.MainController
}


func (this *SettingsController) Index() {
	this.Layout = "admin/index.tpl"


	settings := models.GetSettingsMap()


	this.Data["Title"] = "Website Settings"
	this.Data["Settings"] = settings

	this.TplName = "admin/settings.tpl"
}

func (this *SettingsController) Save(){
	this.Ctx.Request.ParseForm()
	postForm  :=  this.Ctx.Request.PostForm
	o := orm.NewOrm()
	for k, val  := range postForm {
		log.Println(val)
		if created, id, err := o.ReadOrCreate(&models.Setting{Key: k, Value:val[0] }, "Key"); err == nil {
			if created {
				flash := beego.NewFlash()
				flash.Notice("Settings saved")
				flash.Store(&this.Controller)
			} else {
				user := models.Setting{Id: int(id)}
				if o.Read(&user)  == nil {
					user.Key = k
					user.Value = val[0]
					if _, err := o.Update(&user); err == nil {
						flash := beego.NewFlash()
						flash.Notice("Settings saved")
						flash.Store(&this.Controller)
					}
				}
			}
		}
	}
	this.Redirect("/admin/settings/", 302)
}
