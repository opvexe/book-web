package main

import (
	"github.com/astaxie/beego"
	_ "sass-book-web/routers"
	_ "sass-book-web/sysinit"
)

func main() {
	beego.Run()
}
