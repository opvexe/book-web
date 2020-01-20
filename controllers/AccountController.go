package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"sass-book-web/common"
	"sass-book-web/models"
	"sass-book-web/utils"
	"strings"
	"time"
)

type AccountController struct {
	BaseController
}

//注册
func (c *AccountController)Regist()  {
	var (
		nickname string //昵称
		avatar string //头像url
		email string //邮箱
		username string //用户名
		id interface{} //用户id
		captchaOn bool //是否开启验证码
	)
	//如果开启的验证码
	if v ,ok := c.Option["ENABLED_CAPTCHA"];ok &&strings.EqualFold(v,"true"){
		captchaOn = true
		c.Data["CaptchaOn"] = captchaOn
	}
	c.Data["Nickname"] = nickname
	c.Data["Avatar"] =avatar
	c.Data["Email"] = email
	c.Data["Username"] = username
	c.Data["Id"] = id
	c.Data["RandomStr"] = time.Now().Unix()
	c.SetSession("auth",fmt.Sprintf("%v-%v","email",id)) //
	c.TplName = "account/bind.html"
}

//登录
func (c *AccountController) Login()  {
	var member RememberCookie
	//验证cookie
	if cookie,ok := c.GetSecureCookie(common.GetAppKey(),"login");ok {
		if err := utils.Decode(cookie, &member); err == nil {
			if err = c.login(member.MemberId); err == nil {
				c.Redirect(beego.URLFor("HomeController.Index"), 302)
				return
			}
		}
	}
}

func (c *AccountController) login(memberId int) (err error) {
	member,err:=models.FindMemberById(memberId)
	if member.MemberId == 0{
		return errors.New("用户不存在")
	}
	if err!=nil {
		return err
	}
	member.LastLoginTime = time.Now()

}