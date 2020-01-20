package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"sass-book-web/models"
)

type HomeController struct {
	BaseController
}

func (h *HomeController) Index() {
	if cates,err := models.GetCates(-1,1);err!=nil{
		h.Data["Cates"] = cates
	}else {
		beego.Error(errors.New("获取首页分类失败"))
	}
	h.TplName = "home/list.html"
}
