package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)




func init(){
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Post))


	orm.RegisterDataBase("default", "mysql", "root:fusion87@/goblog?charset=utf8")
}
