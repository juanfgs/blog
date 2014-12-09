package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/lib/auth"
)

type User struct {
	Id int
	Username string `orm:"size(100);unique"`
	Password string `orm:"size(100)"`
	Salt string `orm:"size(10)"`


}

type Authenticable interface{
	CheckPassword(password string) bool
	Password(password string) 
}

func (this *User) TableName() string{
	return "users"
}


func (this *User) CheckPassword(password string) bool {
	if this.Password == auth.EncodePassword(password, this.Salt) {
		return true
	}
	return false
}


func (this *User) SetPassword(password string) {
	this.Salt = auth.GetRandomSalt(10)
	this.Password = auth.EncodePassword(password, this.Salt)
}

func  GetUserByName(username string) (user User, err error){
	o := orm.NewOrm()
	err = o.QueryTable("users").Filter("username",username).One(&user)
	return
}

func NewUser (id int) (user User, err error){
	o := orm.NewOrm()
	err = o.QueryTable("users").Filter("id",id).One(&user)
	return
}

