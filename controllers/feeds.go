package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"github.com/juanfgs/rss"
	"github.com/juanfgs/blog/helpers"
	"log"
	"fmt"
)

type FeedsController struct {
	MainController
}

func (this *FeedsController) Index() {


	var posts []models.Post
	o := orm.NewOrm()
	postsPerPage, err := beego.AppConfig.Int("postsperpage")

	if err != nil {
		log.Println(err)
	}

	o.QueryTable("posts").Filter("published", 1).Limit(postsPerPage, 0).OrderBy("-created_at").All(&posts)
	channel := rss.NewChannel("Juan Giménez Silva Posts", "/feeds.xml", "Feed of posts from Juan's blog")

	for _, post := range posts {
		postContent := helpers.RenderPost(post.Content, post.ContentType)
		channel.Add(  post.Title, fmt.Sprintf("/posts/view/%d", post.Id), postContent  )
	}


	content := channel.Marshal()
	this.Ctx.Output.Body( []byte(`<?xml version="1.0" encoding="UTF-8"?><rss version="2.0">`  + string(content) + `</rss>` ) )

	this.ServeXml()
}
