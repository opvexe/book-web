package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

//获取首页数据
func HomeData(pageIndex,pageSize int,cid int,fields ...string) (book []Book,total int,err error) {
	if len(fields) == 0{
		fields = append(fields,"book_id","book_name","identify","cover","order_index")
	}
	fieldStr := "b." +strings.Join(fields,",b.")
	sqlFmt := "select %v from " +TNBook() + "b left join " +TNBookCategory() + "c on b.book_id=c.book_id where c.category_id=" +strconv.Itoa(cid)
	sql := fmt.Sprintf(sqlFmt,fieldStr)
	sqlCount := fmt.Sprintf(sqlFmt, "count(*) cnt")
	fmt.Println("HomeData SQL:",sql)
	o := orm.NewOrm()
	var params []orm.Params
	if _,err := o.Raw(sqlCount).Values(&params);err!=nil{
		if len(params) >0 {
			total, _ = strconv.Atoi(params[0]["cnt"].(string))
		}
	}
	_, err = o.Raw(sql).QueryRows(&book)
	return
}
