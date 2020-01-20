package models

import (
	"github.com/astaxie/beego/orm"
	"sass-book-web/common"
)

//查询
func FindMemberById(id int) (member *Member, err error) {
	member = new(Member) //初始化
	member.MemberId = id
	err = orm.NewOrm().Read(member)
	member.RoleName = common.Role(member.Role)
	return
}
