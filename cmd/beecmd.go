package cmd

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "github.com/juanfgs/blog/models"
)

func main() {


	orm.RegisterDataBase("default", "mysql", "root:fusion87@/goblog?charset=utf8")

	orm.RunCommand()
}
