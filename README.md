#### 1.1 iBook项目

创建项目:

```shell
$ bee new ibook
```

配置数据库:

```shell
db_w_host = 127.0.0.1
db_w_port = 3306
db_w_username = root
db_w_password = 123456
db_w_database = book
```

初始化数据库:

```go
//数据库名称
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	//连接数据库
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
```

创建路由接口:

```go
	//首页&分类&详情
	beego.Router("/", new(controllers.HomeController), "get:Index")
	beego.Router("/explore", new(controllers.ExploreController), "*:Index")
	beego.Router("/book/:key", new(controllers.DocumentController), "*:Index")

	//读书
	beego.Router("/read/:key/:id", new(controllers.DocumentController), "*:Read")
	beego.Router("/read/:key/search", new(controllers.DocumentController), "post:Search")

	//搜索
	beego.Router("/search", new(controllers.SearchController), "get:Search")
	beego.Router("/search/result", new(controllers.SearchController), "get:Result")
```

