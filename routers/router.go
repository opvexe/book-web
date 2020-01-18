package routers

import (
	"github.com/astaxie/beego"
	"sass-book-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
