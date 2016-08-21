package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)




func init(){
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Comment))
	orm.RegisterModel(new(Media))
	orm.RegisterModel(new(Setting))
	orm.RegisterModel(new(Page))
	orm.RegisterModel(new(Notification))
	
	dbpass := beego.AppConfig.String("dbpass")
	dbuser := beego.AppConfig.String("dbuser")
	dbhost := beego.AppConfig.String("dbhost")
	dbname := beego.AppConfig.String("dbname")
	dbcharset := beego.AppConfig.String("dbcharset")
	

	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpass+"@"+dbhost+"/"+dbname+ "?charset=" +dbcharset)
}
