package controllers

type HomeController struct {
	BaseController
}

func (h *HomeController) Index() {
	h.TplName = "home/list.html"
}

func (h *HomeController) Index2() {
	h.TplName = "home/list.html"
}
