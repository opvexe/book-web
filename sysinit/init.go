package sysinit

import _ "sass-book-web/models"

//在main函数调用之前只会调用一次
func init() {
	sysinit()
	dbinit()             //初始化主库
	dbinit("r")          //初始化从库
	dbinit("uaw", "uar") //初始化社区库
}
