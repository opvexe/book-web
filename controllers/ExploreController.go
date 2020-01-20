package controllers

import (
	"github.com/astaxie/beego"
	"math"
	"sass-book-web/models"
	"sass-book-web/utils"
	"strconv"
)

type ExploreController struct {
	BaseController
}

func (t *ExploreController) Index() {

	var (
		cid int //分类id
		cate models.Category
		urlPrefix = beego.URLFor("ExploreController.Index")
	)
	//查询分类
	if cid,_ =t.GetInt("cid");cid>0{
		cate ,_ =models.Find(cid)
		t.Data["Cate"] = cate
	}
	t.Data["Cid"] = cid
	t.TplName = "explore/index.html"

	//获取分页
	pageIndex,_:=t.GetInt("page",1)
	pageSize := 24

	books,total,err:=models.HomeData(pageIndex,pageSize,cid)
	if err!=nil {
		beego.Error(err)
		t.Abort("404")
	}
	if total>0{
		urlSuffix := ""
		if cid>0{
			urlSuffix = urlSuffix +"&cid=" +strconv.Itoa(cid)
		}
		html := utils.NewPaginations(4, total, pageSize, pageIndex, urlPrefix, urlSuffix)
		t.Data["PageHtml"] = html
	}else {
		t.Data["PageHtml"] = ""
	}
	t.Data["TotalPages"] = int(math.Ceil(float64(total)/float64(pageSize)))
	t.Data["List"] = books
}