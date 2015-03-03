package controllers 

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/juanfgs/blog/models"
	"time"
	"log"
	"net/http"
	"net/url"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
)


type CommentsController struct {
	MainController
}

type ReCaptchaResponse struct{
	Success bool `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

func (this *CommentsController) CommentWrite(){
	var postid int
	var err error
	flash := beego.NewFlash()
	o := orm.NewOrm()
	
	comment := new(models.Comment)
	comment.Commenter = this.GetString("commenter")
	comment.Comment = this.GetString("comment")
	var recaptcha_response = this.GetString("g-recaptcha-response")
	recaptchaKey := beego.AppConfig.String("recaptchaKey")

	recaptchaData := url.Values{}
	recaptchaData.Set("secret", recaptchaKey)
	recaptchaData.Set("response", recaptcha_response)
	client := &http.Client{}
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest("POST","https://www.google.com/recaptcha/api/siteverify", bytes.NewBufferString(recaptchaData.Encode()))
	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(recaptchaData.Encode())))
	res, err = client.Do(req)
	var jsonbytes []byte
	jsonbytes, err = ioutil.ReadAll(res.Body)
	var recaptcharesponse ReCaptchaResponse
	err = json.Unmarshal(jsonbytes,&recaptcharesponse)
	if err != nil {
		log.Println(err)
	}

	if recaptcharesponse.Success {
		

		if postid, err = this.GetInt("post_id"); err == nil{
			var post models.Post
			o.QueryTable("posts").Filter("id", postid).One(&post)
			comment.Post = &post
		}
		comment.CreatedAt = time.Now()
		comment.UpdatedAt = time.Now()
		_,err = o.Insert(comment)
		if err != nil {
			log.Println(err)
			flash.Error("Error inserting Comment")
			flash.Store(&this.Controller)
			return 
		}

		flash.Notice("Comment Added")
		flash.Store(&this.Controller)
		this.Redirect(this.Ctx.Input.Refer(), 302)
	} else {
		flash.Notice("Captcha failed")
		flash.Store(&this.Controller)
		log.Println(recaptcharesponse)
		this.Redirect(this.Ctx.Input.Refer(), 302)
	}
}
