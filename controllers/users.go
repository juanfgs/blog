package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/juanfgs/blog/models"
	"log"
)

type UsersController struct {
	MainController
}

func (this *UsersController) Login() {
	this.Layout = "admin/index.tpl"

	this.TplNames = "login.tpl"
}

func (this *UsersController) LoginPost() {

	username := this.GetString("Username")
	password := this.GetString("Password")

	if user, err := models.GetUserByName(username); err == nil {
		if user.CheckPassword(password) {
			var sessionName = beego.AppConfig.String("SessionName")
			
			v := this.GetSession(sessionName)

			if v == nil {
				this.SetSession(sessionName, user.Id)
			}		
			this.Redirect("/", 302)
		}
	} else {
		log.Println(err)
	}

	this.Redirect("/login", 302)
}

func (this *UsersController) Logout() {

}

func (this *UsersController) RegisterPost() {

	valid := validation.Validation{}
	username := this.GetString("Username")
	password := this.GetString("Password")
	valid.Required(username, "Username")
	valid.Required(password, "Password")
	if valid.HasErrors() {
		flash := beego.NewFlash()
		for _, err := range valid.Errors {
			flash.Error(err.Key, err.Message)
			log.Println(err)
		}

	} else {
		o := orm.NewOrm()

		user := new(models.User)
		user.Username = username

		user.SetPassword(password)
		_, err := o.Insert(user)
		if err != nil {
			log.Println(err)
		}

	}
	this.Redirect("/", 302)

}

func (this *UsersController) URLMapping() {
	this.Mapping("Login", this.Login)
	this.Mapping("LoginPost", this.LoginPost)
	this.Mapping("Logout", this.Logout)
	this.Mapping("RegisterPost", this.RegisterPost)

}
