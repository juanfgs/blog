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

	this.TplName = "login.tpl"
}

func (this *UsersController) LoginPost() {
	username := this.GetString("Username")
	password := this.GetString("Password")
	flash := beego.NewFlash()
	if user, err := models.GetUserByName(username); err == nil {
		if user.CheckPassword(password) {
			var sessionName = beego.AppConfig.String("SessionName")

			v := this.GetSession(sessionName)

			if v == nil {
				this.SetSession(sessionName, user.Id)

			}
			flash.Notice("Login Successful")
			flash.Store(&this.Controller)
			this.Redirect("/admin/", 302)
			return
		} else {

			flash.Error("Invalid Password")
			flash.Store(&this.Controller)
			this.Redirect("/login", 302)
			return
		}
	} else {
		flash.Error( "User doesn't exist")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
		log.Println(err)
	}

	this.Redirect("/admin/", 302)
}

func (this *UsersController) Logout() {

}

func (this *UsersController) RegisterPost() {
	o := orm.NewOrm()

	userscount, _ := o.QueryTable("users").Count()
	flash := beego.NewFlash()
	if userscount > 0 && beego.AppConfig.String("SingleUser") == "yes" {
		flash.Error("Single user mode enabled and administrator already exists")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
	}


	valid := validation.Validation{}
	username := this.GetString("Username")


	password := this.GetString("Password")
	valid.Required(username, "Username")
	valid.Required(password, "Password")

	if o.QueryTable("users").Filter("username", username).Exist() {
		flash.Error( "User already exists")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			flash.Error(err.Key, err.Message)
			flash.Store(&this.Controller)
			this.Redirect("/login", 302)
			return
		}
	} else {

		user := new(models.User)
		user.Username = username

		user.SetPassword(password)
		_, err := o.Insert(user)
		if err != nil {
			log.Println(err)
		}
		flash.Notice("User created successfully")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
	}
	this.Redirect("/login", 302)
	return
}

func (this *UsersController) URLMapping() {
	this.Mapping("Login", this.Login)
	this.Mapping("LoginPost", this.LoginPost)
	this.Mapping("Logout", this.Logout)
	this.Mapping("RegisterPost", this.RegisterPost)

}
