package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"
	"io"
	"os"
	"strconv"
	"time"
)

type MediaController struct {
	controllers.MainController
}

func (this *MediaController) Create() {

	if file, header, err := this.GetFile("media"); err != nil {
		this.Abort("500")

	} else {
		defer file.Close()

		secs := strconv.Itoa(int(time.Now().Unix()))

		newFileName := secs + "_" + header.Filename
		UploadsDir := beego.AppConfig.String("uploads_dir")

		destFile, err := os.OpenFile(UploadsDir+newFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			this.Abort("500")
		}
		io.Copy(destFile, file)
		defer destFile.Close()

		o := orm.NewOrm()
		media := new(models.Media)
		media.Type = header.Header["Content-Type"][0]
		media.Filename = newFileName
		media.CreatedAt = time.Now()
		media.UpdatedAt = time.Now()
		_, err = o.Insert(media)

		if err != nil {
			this.Abort("500")
		}
		this.Data["json"] = &media
		this.ServeJson()

	}

}

func (this *MediaController) Upload() {
	this.Layout = "admin/index.tpl"
	this.Data["title"] = "Upload media"

	this.TplNames = "admin/uploadform.tpl"
}

func (this *MediaController) Index() {
	this.Layout = "admin/index.tpl"

	var media []models.Media
	o := orm.NewOrm()
	mediaPerPage := 10
	countMedia, err := o.QueryTable("media").Count()
	if err != nil {
		this.Abort("500")
	}
	paginator := pagination.SetPaginator(this.Ctx, mediaPerPage, countMedia)
	o.QueryTable("media").Limit(mediaPerPage, paginator.Offset()).OrderBy("-created_at").All(&media)

	this.Data["media"] = media

	this.TplNames = "admin/media.tpl"

}

func (this *MediaController) Delete() {
	mediaId, err := this.GetInt(":id")
	UploadsDir := beego.AppConfig.String("uploads_dir")
	if err != nil {
		this.Abort("400")
	}
	o := orm.NewOrm()

	media := new(models.Media)
	o.QueryTable("media").Filter("id", mediaId).One(media)
	err = os.Remove(UploadsDir + media.Filename)
	if _, err = o.Delete(media); err != nil {
		this.Abort("500")
	}

	this.Redirect("/admin/media/", 302)

}
