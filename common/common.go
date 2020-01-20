package common

import "github.com/astaxie/beego"

const (
	WorkingDirectory = "./"
	SessionName      = "sass-book-session"
)

//用户权限
const (
	MemberSuperRole   = 0 //超级管理员
	MemberAdminRole   = 1 //普通管理员
	MemberGeneralRole = 2 //普通用户
)

//获取管理员权限
func Role(role int) string {
	switch role {
	case MemberSuperRole:
		return "超级管理员"
	case MemberAdminRole:
		return "管理员"
	case MemberGeneralRole:
		return "普通用户"
	default:
		return ""
	}
}

//获取cookie的key
func GetAppKey() string {
	return beego.AppConfig.DefaultString("app_key", "godoc")
}
