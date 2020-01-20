package models

import "time"

// 书籍表
type Book struct {
	BookId         int       `orm:"pk;auto" json:"book_id"`           //主键
	BookName       string    `orm:"size(500)" json:"book_name"`       //书名
	Identify       string    `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int       `orm:"default(0)" json:"order_index"`
	Description    string    `orm:"size(1000)" json:"description"`     //图书描述
	Cover          string    `orm:"size(1000)" json:"cover"`           //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`            //编辑器类型:markdown
	Status         string    `orm:"default(0)" json:"status"`          //状态:0:正常 1删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"` //是否私有 0：公开 1私有
	PrivateToken   string    `orm:"size(500)" json:"private_token"`    //生成图书私有访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	DocCount       int       `json:"doc_count"`                          //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`      //评论数量
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int       `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                                       //作者
	AuthorURL      string    `orm:"column(author_url);size(1000)"`                  //来源连接
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"` //修改时间
	ReleaseTime    time.Time `orm:"type(datetime)" json:"release_time"`             //发布时间
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

//图书收藏表
type Collection struct {
	Id       int `orm:"pk;auto"`
	MemberId int `orm:"index"`
	BookId   int
}

//成员表
type Member struct {
	MemberId      int       `orm:"pk;auto" json:"member_id"`
	Account       string    `orm:"size(30);unique" json:"account"`
	Nickname      string    `orm:"size(30);unique" json:"nickname"`
	Password      string    `json:"-"`
	Description   string    `orm:"size(400)" json:"description"`
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

//评论表
type Comments struct {
	Id         int       `orm:"pk;auto"`
	Uid        int       `orm:"index"` //用户id
	BookId     int       `orm:"index"` //文档项目id
	Content    string    //评价内容
	TimeCreate time.Time //评论时间
}

//评分表
type Score struct {
	Id         int `orm:"pk;auto"`
	BookId     int
	Uid        int
	Score      int
	TimeCreate time.Time //评分时间
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

//附件表
type Attachment struct {
	AttachmentId int `orm:"pk;auto" json:"attachment_id"`
	BookId       int `json:"book_id"`
	DocumentId   int `json:"doc_id"`
	Name         string
	Path         string    `orm:"size(2000)" json:"file_path"`
	Size         float64   `orm:"type(float)" json:"file_size"`
	Ext          string    `orm:"size(50)" json:"file_ext"`
	HttpPath     string    `orm:"size(2000)" json:"http_path"`
	CreateTime   time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	CreateAt     int       `orm:"type(int)" json:"create_at"`
}

//文档编辑
type DocumentStore struct {
	DocumentId int    `orm:"pk;auto;column(document_id)"`
	Markdown   string `orm:"type(text)"` //markdown内容
	Content    string `orm:"type(text)"` //html内容
}

//粉丝表
type Fans struct {
	Id       int `orm:"pk;auto"`
	MemberId int
	FansId   int `orm:"index"` //粉丝id
}

type Relationship struct {
	RelationshipId int `orm:"pk;auto" json:"relationship_id"`
	MemberId       int `json:"member_id"`
	BookId         int ` json:"book_id"`
	RoleId         int `json:"role_id"` // common.BookRole
}

//自定义表名
func (c *Book) TableName() string {
	return TNBook()
}

func (c *Category) TableName() string {
	return TNCategory()
}

func (c *BookCategory) TableName() string {
	return TNBookCategory()
}

func (c *Document) TableName() string {
	return TNDocuments()
}

func (c *DocumentStore) TableName() string {
	return TNDocumentStore()
}

func (c *Attachment) TableName() string {
	return TNAttachment()
}

func (c *Relationship) TableName() string {
	return TNRelationship()
}

func (c *Member) TableName() string {
	return TNMembers()
}

func (c *Collection) TableName() string {
	return TNCollection()
}

func (c *Fans) TableName() string {
	return TNFans()
}

func (c *Comments) TableName() string {
	return TNComments()
}

func (c *Score) TableName() string {
	return TNScore()
}
