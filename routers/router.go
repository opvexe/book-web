package routers

import (
	"github.com/astaxie/beego"
	"sass-book-web/controllers"
)

func init() {

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

	//登录
	beego.Router("/login", new(controllers.AccountController), "*:Login")
	beego.Router("/regist", new(controllers.AccountController), "*:Regist")
	beego.Router("/logout", new(controllers.AccountController), "*:Logout")
	beego.Router("/doregist", new(controllers.AccountController), "post:DoRegist")

	//编辑
	beego.Router("/api/:key/edit/?:id", new(controllers.DocumentController), "*:Edit")
	beego.Router("api/:key/content/?:id", new(controllers.DocumentController), "*:Content")
	beego.Router("api/upload", new(controllers.DocumentController), "post:Upload")
	beego.Router("/api/:key/create", new(controllers.DocumentController), "post:Create")
	beego.Router("/api/:key/delete", new(controllers.DocumentController), "post:Delete")

	//用户图书管理
	beego.Router("/book", new(controllers.BookController), "*Index")                          //我的图书
	beego.Router("/book/create", new(controllers.BookController), "post:Create")              //创建图书
	beego.Router("/book/:key/setting", new(controllers.BookController), "*:Setting")          //图书设置
	beego.Router("/book/setting/upload", new(controllers.BookController), "post:UploadCover") //图书封面
	beego.Router("/book/star/:id", new(controllers.BookController), "*:Collection")           //收藏图书
	beego.Router("/book/setting/save", new(controllers.BookController), "post:SaveBook")      //保存
	beego.Router("/book/:key/release", new(controllers.BookController), "post:Release")       //发布
	beego.Router("/book/setting/token", new(controllers.BookController), "post:CreateToken")  //创建Token

	//个人中心
	beego.Router("/user/:username", new(controllers.UserController), "get:Index")                 //分享
	beego.Router("/user/:username/collection", new(controllers.UserController), "get:Collection") //收藏
	beego.Router("/user/:username/follow", new(controllers.UserController), "get:Follow")         //关注
	beego.Router("/user/:username/fans", new(controllers.UserController), "get:Fans")             //粉丝
	beego.Router("/follow/:uid", new(controllers.UserController), "get:SetFollow")                //关注或取消关注
	beego.Router("/book/score/:id", new(controllers.UserController), "*:Score")                   //评分
	beego.Router("/book/comment/:id", new(controllers.UserController), "post:Comment")            //评论

	//个人设置
	beego.Router("/setting", new(controllers.SettingController), "*:Index")
	beego.Router("/setting/upload", new(controllers.SettingController), "*:Upload")

	//管理后台
	beego.Router("/manager/category", new(controllers.ManagerController), "post,get:Category")
	beego.Router("/manager/update-cate", new(controllers.ManagerController), "get:UpdateCate")
	beego.Router("/manager/del-cate", new(controllers.ManagerController), "get:DelCate")
	beego.Router("/manager/icon-cate", new(controllers.ManagerController), "post:UpdateCateIcon")
}
