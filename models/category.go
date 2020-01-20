package models

import "github.com/astaxie/beego/orm"

// 查询所有分类
// pid 分类id
//status 显示状态
//对应的sql语句:select * from md_category where pid=1 and where status = 0 order by status desc,sort asc ,title asc
func GetCates(pid int,status int) (cates []Category,err error) {
	qs := orm.NewOrm().QueryTable(TNCategory())
	if pid>-1{
		qs = qs.Filter("pid",pid)
	}
	if status == 0|| status == 1 {
		qs = qs.Filter("status",status)
	}
	_,err=qs.OrderBy("-status","sort","title").All(&cates)
	return
}

//查询分类 [根据分类id]
//cid 分类id
//sql语句：select *from md_category where id = cid
func Find(cid int) (cate Category,err error) {
	cate.Id = cid
	err = orm.NewOrm().Read(&cate)
	return
}
