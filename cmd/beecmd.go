package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "github.com/juanfgs/blog/models"
)

func main() {

	dbpass := beego.AppConfig.String("dbpass")
	dbuser := beego.AppConfig.String("dbuser")
	dbhost := beego.AppConfig.String("dbhost")
	dbname := beego.AppConfig.String("dbname")
	dbport := beego.AppConfig.String("dbport")
	dbcharset := beego.AppConfig.String("dbcharset")

	orm.RegisterDataBase("cmd", "mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname+ "?charset=" +dbcharset)

	orm.RunCommand()
}
