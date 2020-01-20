package controllers

type DocumentController struct {
	BaseController
}

//图书目录&详情页
func (c *DocumentController) Index() {
	//token := c.GetString("token")
	//identify := c.Ctx.Input.Param(":key")
	//if identify == "" {
	//	c.Abort("404")
	//}
	//tab := strings.ToLower(c.GetString("tab"))

}

//判断图书内容并判断权限
func (t *DocumentController) getBookData(identify, token string) {

}
