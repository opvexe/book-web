package sysinit

import (
	"github.com/astaxie/beego"
	"log"
	conf "sass-book-web/common"
	"path/filepath"
	"strings"
)

func sysinit() {
	uploads := filepath.Join(conf.WorkingDirectory, "uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads
	//注册前端使用函数
	registerFunctions()
}

func registerFunctions() {
	if err :=beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
			return cdn + "/" + p
		}
		return cdn + p
	});err!=nil{
		log.Printf("AddFuncMap:%s",err)
	}
}
