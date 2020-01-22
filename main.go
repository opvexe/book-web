package main

import (
	"fmt"
	_ "ziyoubiancheng/mbook/routers"
	_ "ziyoubiancheng/mbook/sysinit"
	"ziyoubiancheng/mbook/utils/pagecache"

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
