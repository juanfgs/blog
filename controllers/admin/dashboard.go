package admin

import(
	"github.com/juanfgs/blog/controllers"
)

type DashboardController struct {
	controllers.MainController
}

func (this *DashboardController) Index() {
	this.Layout = "admin/index.tpl"
	this.Data["Title"] = "Dashboard"

	this.TplName = "admin/dashboard.tpl"
}


