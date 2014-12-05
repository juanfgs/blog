package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init(){
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Category))
	orm.RegisterDataBase("default", "mysql", "root:fusion87@/goblog?charset=utf8")
}
