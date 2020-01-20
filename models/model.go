package models

import "time"

// 书籍表
type Book struct {
	BookId         int    `orm:"pk:auto" json:"book_id"`           //主键
	BookName       string `orm:"size(500)" json:"book_name"`       //书名
	Identify       string `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int    `orm:"default(0)" json:"order_index"`
	Description    string `orm:"size(1000)" json:"description"`     //图书描述
	Cover          string `orm:"size(1000)" json:"cover"`           //封面地址
	Editor         string `orm:"size(50)" json:"editor"`            //编辑器类型:markdown
	Status         string `orm:"default(0)" json:"status"`          //状态:0:正常 1删除
	PrivatelyOwned int    `orm:"default(0)" json:"privately_owned"` //是否私有 0：公开 1私有
	PrivateToken   string `orm:"size(500)" json:"private_token"`    //生成图书私有访问Token
	MemberId       int    `orm:"size(100)" json:"member_id"`
	DocCount       int    `json:"doc_count"`                          //文档数量
	CommentCount   int    `orm:"type(int)" json:"comment_count"`      //评论数量
	Vcnt           int    `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int    `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int    `orm:"default(40)" json:"score"`            //评分
	CntScore       int    //评分人数
	CntComment     int    //评论人数
	Author         string `orm:"size(50)"`                      //作者
	AuthorURL      string `orm:"column(author_url);size(1000)"` //来源连接
	CreateTime  time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime  time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"` //修改时间
	ReleaseTime time.Time `orm:"type(datetime)" json:"release_time"`             //发布时间
	//auto_now 每次 model 保存时都会对时间自动更新
	//auto_now_add 第一次保存时才设置时间
}

//书籍分类表
type Category struct {
	Id     int    `orm:"pk;auto"`
	Pid    int    `orm:"default(0)"` //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string `orm:"size(256)"` //说明
	Icon   string
	Cnt    int  //统计分类下的图书
	Sort   int  //排序
	Status bool //状态,true显示
}

//图书分类对应的关系表
type BookCategory struct {
	Id         int `orm:"pk;auto"`
	BookId     int //书籍id
	CategoryId int //分类id
}

//自定义表名
func (b *Book) TableName() string {
	return TNBook()
}

func (c *Category) TableName() string {
	return TNCategory()
}

func (bc *BookCategory) TableName() string {
	return TNBookCategory()
}
