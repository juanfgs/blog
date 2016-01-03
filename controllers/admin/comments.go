package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/juanfgs/blog/controllers"
	"github.com/astaxie/beego/utils/pagination"
	"log"
)

type CommentsController struct {
	controllers.MainController
}


func (this *CommentsController) Index() {
	this.Layout = "admin/index.tpl"

	var comments []models.Comment
	o := orm.NewOrm()
	commentsPerPage := 10
	countComments, err := o.QueryTable("comments").Count()
	if err != nil {
		log.Println(err)
	}

	paginator := pagination.SetPaginator(this.Ctx, commentsPerPage, countComments)

	o.QueryTable("comments").Limit(commentsPerPage, paginator.Offset()).OrderBy("-created_at").All(&comments)

	this.Data["Title"] = "Listing comments"
	this.Data["comments"] = comments
	this.TplNames = "admin/comments/index.tpl"
}

func (this *CommentsController) Delete(){

	commentId, err:= this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	comment := new(models.Comment)
	o.QueryTable("comments").Filter("id", commentId).One(comment)
	if _, err = o.Delete(comment); err == nil {
		flash.Notice("Comment erased")
		flash.Store(&this.Controller)
	} else {
		flash.Notice("Cannot delete comment, check log for details")
		log.Println(err)
		flash.Store(&this.Controller)
	}
	this.Redirect("/admin/comments/", 302)
	return
}

func (this *CommentsController) Edit() {
	this.Layout = "admin/index.tpl"
	commentId, err := this.GetInt(":id")

	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	comment := new(models.Comment)


	o.QueryTable("comments").Filter("id", commentId).One(comment)
	this.Data["Title"] = "Editing " + comment.Commenter + "'s comment"
	this.Data["Commenter"] = comment.Commenter
	this.Data["Comment"] = comment.Comment

	this.TplNames = "admin/comments/edit.tpl"
}

func (this *CommentsController) EditWrite() {
	commentId, err := this.GetInt(":id")
	flash := beego.NewFlash()
	if err != nil {
		this.Abort("400")
	}

	comment := new(models.Comment)

	o := orm.NewOrm()

	o.QueryTable("comments").Filter("id", commentId).One(comment)

	if val := this.GetString("Commenter"); val != comment.Commenter {
		comment.Commenter = val
	}
	if val := this.GetString("Comment"); val != comment.Comment {
		comment.Comment = val
	}

	if _, err := o.Update(comment); err == nil {
		flash.Notice("Comment Saved")
		flash.Store(&this.Controller)
		this.Redirect("/admin/comments/", 302)
		return
	} else {

		this.Abort("500")
	}

}
