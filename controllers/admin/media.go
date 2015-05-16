package admin

import (
	"github.com/astaxie/beego"
	"github.com/juanfgs/blog/controllers"
	"github.com/juanfgs/blog/models"	
	"os"
	"io"
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	
)

type MediaController struct {
	controllers.MainController
}

func (this *MediaController) Create() {
	
	if file, header, err := this.GetFile("Media"); err != nil {
		this.Abort("500")
		
	} else {
		defer file.Close()

		secs := strconv.Itoa( int(time.Now().Unix()) )
		uploads_dir := beego.AppConfig.String("uploads_dir")
		newFileName := secs + "_" + header.Filename
		
		destFile, err := os.OpenFile(uploads_dir + newFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			this.Abort("500")
		}
		io.Copy(destFile, file)
		defer destFile.Close()

		o := orm.NewOrm()
		media := new(models.Media)

		media.Filename = newFileName
		media.CreatedAt = time.Now()
		media.UpdatedAt = time.Now()
		_,err = o.Insert(media)

		if err != nil {
			this.Abort("500")
		}

	}
	

	

	
	this.Redirect("/admin/", 302)

}

func (this *MediaController) Upload() {
	this.Layout = "admin/index.tpl"
	this.Data["title"] = "Upload media"

	this.TplNames = "admin/uploadform.tpl"
}


