package controllers

import (
	"github.com/astaxie/beego"
	"sass-book-web/common"
	"sass-book-web/models"
	"sass-book-web/utils"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	Member          *models.Member    //用户
	Option          map[string]string //全局设置
	EnableAnonymous bool              //开启匿名访问
}

//cookie
type RememberCookie struct {
	MemberId int
	Account  string
	Time     time.Time
}

//每个子类Controller公用方法调用前，都执行一下Prepare方法
func (c *BaseController) Prepare() {
	c.Member = new(models.Member)
	c.EnableAnonymous = false

	//从session中获取用户信息
	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.MemberId > 0 {
		c.Member = &member
	} else {
		//如果cookie中存在登录信息，从cookie中获取用户信息
		if cookie, ok := c.GetSecureCookie(common.GetAppKey(), "login"); ok {
			var remember RememberCookie
			err := utils.Decode(cookie, &remember)
			if err == nil {
				member, err := models.FindMemberById(remember.MemberId)
				if err == nil {
					c.SetMember(*member)
				}
			}
		}
	}
	if c.Member.RoleName == "" {
		c.Member.RoleName = common.Role(c.Member.MemberId)
	}
	c.Data["Member"] = c.Member
	c.Data["BaseUrl"] = c.BaseURL()
	c.Data["SITE_NAME"] = "FaceBook"
	//设置全局配置
	c.Option = make(map[string]string)
	c.Option["ENABLED_CAPTCHA"] = "false"
}

//设置用户登录信息
func (c *BaseController) SetMember(member models.Member) {
	if member.MemberId <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(common.SessionName, member)
		c.SetSession("uid", member.MemberId)
	}
}

//设置请求路径
func (c *BaseController) BaseURL() string {
	host := beego.AppConfig.String("sitemap_host")
	if len(host) > 0 {
		if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
			return host
		}
		return c.Ctx.Input.Scheme() + "://" + host
	}
	return c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
}
