package main

import (
	"github.com/astaxie/beego"
	_ "sass-book-web/sysinit"
	_ "sass-book-web/routers"
)

func main() {
	beego.Run()
}
