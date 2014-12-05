package main

import (
	"github.com/astaxie/beego"

	_ "github.com/juanfgs/blog/models"
	_ "github.com/juanfgs/blog/routers"
)

func main() {
	beego.Run()
}
