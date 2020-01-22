package main

import (
	"fmt"
	_ "shumin-project/sass-book-web/routers"
	_ "shumin-project/sass-book-web/sysinit"
	"shumin-project/sass-book-web/utils/pagecache"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	task := toolbox.NewTask("clear_expired_cache", "1 1 2 * * *", func() error { fmt.Println("--delete cache---"); pagecache.ClearExpiredFiles(); return nil })
	toolbox.AddTask("mbook_task", task)
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}
