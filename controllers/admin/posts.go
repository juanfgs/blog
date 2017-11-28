package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"
	"log"
)

type PostsController struct {
	controllers.MainController
}

func (this *PostsController) GetPost(){

	id,err := this.GetInt(":id")
	if err != nil {
		log.Println(id)
		log.Println(err)
		this.Abort("400")
	}

	var post models.Post
	o := orm.NewOrm()
	err = o.QueryTable("posts").Filter("id", id).One(&post)

	this.Data["post"] = &post
	this.ServeJSON()
}

func (this *PostsController) GetPosts(){
	var posts []models.Post
	var err error

	o := orm.NewOrm()
	_, err = o.QueryTable("posts").All(&posts)

	if err != nil{
		log.Println(err)
		this.Abort("400")
	}
	this.Data["json"] = &posts
	this.ServeJSON()

}
