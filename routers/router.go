package routers

import (
	"github.com/astaxie/beego"
	"sass-book-web/controllers"
)

func init() {

	//首页&分类&详情
	beego.Router("/", new(controllers.HomeController),"get:Index")
	beego.Router("/2", new(controllers.HomeController),"get:Index2")

}
