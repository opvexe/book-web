package models

import "time"

//Book表
type Book struct {
	BookId         int       `orm:"pk;auto" json:"book_id"`
	BookName       string    `orm:"size(500)" json:"book_name"`       //名称
	Identify       string    `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int       `orm:"default(0)" json:"order_index"`
	Description    string    `orm:"size(1000)" json:"description"`       //图书描述
	Cover          string    `orm:"size(1000)" json:"cover"`             //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`              //编辑器类型: "markdown"
	Status         int       `orm:"default(0)" json:"status"`            //状态:0 正常 ; 1 已删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"`   // 是否私有: 0 公开 ; 1 私有
	PrivateToken   string    `orm:"size(500);null" json:"private_token"` // 私有图书访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"`
	ReleaseTime    time.Time `orm:"type(datetime);" json:"release_time"` //发布时间
	DocCount       int       `json:"doc_count"`                          //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int       `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                      //来源
	AuthorURL      string    `orm:"column(author_url);size(1000)"` //来源链接
}

//分类表
type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string //介绍
	Icon   string
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

//图书分类对应关系
type BookCategory struct {
	Id         int //自增主键
	BookId     int //书籍id
	CategoryId int //分类id
}

//附件表
type Attachment struct {
	AttachmentId int `orm:"pk;auto" json:"attachment_id"`
	BookId       int ` json:"book_id"`
	DocumentId   int ` json:"doc_id"`
	Name         string
	Path         string    `orm:"size(2000)" json:"file_path"`
	Size         float64   `orm:"type(float)" json:"file_size"`
	Ext          string    `orm:"size(50)" json:"file_ext"`
	HttpPath     string    `orm:"size(2000)" json:"http_path"`
	CreateTime   time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	CreateAt     int       `orm:"type(int)" json:"create_at"`
}

//图书章节内容
type Document struct {
	DocumentId   int           `orm:"pk;auto;column(document_id)" json:"doc_id"`
	DocumentName string        `orm:"column(document_name);size(500)" json:"doc_name"`
	Identify     string        `orm:"column(identify);size(100);index;null;default(null)" json:"identify"`
	BookId       int           `orm:"column(book_id);type(int)" json:"book_id"`
	ParentId     int           `orm:"column(parent_id);type(int);default(0)" json:"parent_id"`
	OrderSort    int           `orm:"column(order_sort);default(0);type(int)" json:"order_sort"`
	Release      string        `orm:"column(release);type(text);null" json:"release"`
	CreateTime   time.Time     `orm:"column(create_time);type(datetime);auto_now_add" json:"create_time"`
	MemberId     int           `orm:"column(member_id);type(int)" json:"member_id"`
	ModifyTime   time.Time     `orm:"column(modify_time);type(datetime);default(null);auto_now" json:"modify_time"`
	ModifyAt     int           `orm:"column(modify_at);type(int)" json:"-"`
	Version      int64         `orm:"type(bigint);column(version)" json:"version"`
	AttachList   []*Attachment `orm:"-" json:"attach"`
	Vcnt         int           `orm:"column(vcnt);default(0)" json:"vcnt"`
	Markdown     string        `orm:"-" json:"markdown"`
}

// 用户表
type Member struct {
	MemberId      int       `orm:"pk;auto" json:"member_id"`
	Account       string    `orm:"size(30);unique" json:"account"`
	Nickname      string    `orm:"size(30);unique" json:"nickname"`
	Password      string    ` json:"-"`
	Description   string    `orm:"size(640)" json:"description"`
	Email         string    `orm:"size(100);unique" json:"email"`
	Phone         string    `orm:"size(20);null;default(null)" json:"phone"`
	Avatar        string    `json:"avatar"`
	Role          int       `orm:"default(1)" json:"role"`
	RoleName      string    `orm:"-" json:"role_name"`
	Status        int       `orm:"default(0)" json:"status"`
	CreateTime    time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	CreateAt      int       `json:"create_at"`
	LastLoginTime time.Time `orm:"type(datetime);null" json:"last_login_time"`
}

//文档编辑
type DocumentStore struct {
	DocumentId int    `orm:"pk;auto;column(document_id)"`
	Markdown   string `orm:"type(text);"` //markdown内容
	Content    string `orm:"type(text);"` //html内容
}

//粉丝表
type Fans struct {
	Id       int //PK
	MemberId int
	FansId   int `orm:"index"` //粉丝id
}

//收藏表
type Collection struct {
	Id       int
	MemberId int `orm:"index"`
	BookId   int
}

//关注表
type Relationship struct {
	RelationshipId int `orm:"pk;auto;" json:"relationship_id"`
	MemberId       int `json:"member_id"`
	BookId         int ` json:"book_id"`
	RoleId         int `json:"role_id"` // common.BookRole
}

//评分表
type Score struct {
	Id         int
	BookId     int
	Uid        int
	Score      int //评分
	TimeCreate time.Time
}

//评论表
type Comments struct {
	Id         int
	Uid        int       `orm:"index"` //用户id
	BookId     int       `orm:"index"` //文档项目id
	Content    string    //评论内容
	TimeCreate time.Time //评论时间
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *BookCategory) TableName() string {
	return TNBookCategory()
}

// 多字段唯一键
func (m *BookCategory) TableUnique() [][]string {
	return [][]string{
		[]string{"BookId", "CategoryId"},
	}
}

func (m *Book) TableName() string {
	return TNBook()
}

func (m *DocumentStore) TableName() string {
	return TNDocumentStore()
}

func (m *Attachment) TableName() string {
	return TNAttachment()
}

func (m *Relationship) TableName() string {
	return TNRelationship()
}

//  联合唯一索引
func (m *Relationship) TableUnique() [][]string {
	return [][]string{
		[]string{"MemberId", "BookId"},
	}
}

func (m *Member) TableName() string {
	return TNMembers()
}

func (m *Fans) TableName() string {
	return TNFans()
}

// 多字段唯一键
func (m *Fans) TableUnique() [][]string {
	return [][]string{
		[]string{"MemberId", "FansId"},
	}
}

func (m *Score) TableName() string {
	return TNScore()
}

// 多字段唯一键
func (m *Score) TableUnique() [][]string {
	return [][]string{
		[]string{"Uid", "BookId"},
	}
}

func (m *Comments) TableName() string {
	return TNComments()
}


func (m *Document) TableUnique() [][]string {
	return [][]string{
		[]string{"BookId", "Identify"},
	}
}

// 设置index
func (m *Document) TableIndex() [][]string {
	return [][]string{
		[]string{"BookId", "ParentId", "OrderSort"},
	}
}

func (m *Collection) TableName() string {
	return TNCollection()
}

// 多字段唯一键
func (m *Collection) TableUnique() [][]string {
	return [][]string{
		[]string{"MemberId", "BookId"},
	}
}