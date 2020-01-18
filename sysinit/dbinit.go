package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

//dbinit("w") 初始化主库
func dbinit(aliases ...string) {
	//如果是主库，自动建表
	isDev := ("dev" == beego.AppConfig.String("runmode"))
	if len(aliases) > 0 {
		for _, aliaas := range aliases {
			if "w" == aliaas {
				if err := orm.RunSyncdb("default", false, true); err != nil {
					log.Fatalf("自动建表:%s", err)
				}
			}
		}
	} else {
		registerDatabases("w")
		//自动建表
		if err := orm.RunSyncdb("default", false, true); err != nil {
			log.Fatalf("自动建表:%s", err)
		}
	}
	if isDev {
		orm.Debug = isDev
	}
}

//注册单数据库
func registerDatabases(alias string) {
	if len(alias) <= 0 {
		return
	}
	//连接名称
	dbAlias := alias //default
	if "w" == alias || "default" == alias || len(alias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}
	//数据库名称
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库用户名
	dbuser := beego.AppConfig.String("db_" + alias + "_username")
	//数据库密码
	dbpwd := beego.AppConfig.String("db_" + alias + "_password")
	//数据库端IP
	dbhost := beego.AppConfig.String("db_" + alias + "_host")
	//数据库端口号
	dbport := beego.AppConfig.String("db_" + alias + "_port")
	//root:123456@tcp(127.0.0.1:3306)/book?charset=utf8
	//连接数据库
	if err := orm.RegisterDataBase(dbAlias, "mysql", dbuser+":"+dbpwd+"@tcp("+dbhost+":"+dbport+")"+dbName+"?charset=utf8", 30); err != nil {
		log.Fatalf("连接数据库失败:%s", err)
	}
}
