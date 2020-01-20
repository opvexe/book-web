package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Category),
		new(Book),
		new(BookCategory),
		new(Document),
		new(DocumentStore),
		new(Attachment),
		new(Member),
		new(Collection),
		new(Relationship),
		new(Fans),
		new(Comments),
		new(Score),
	)
}

//分类表
func TNCategory() string {
	return "md_category"
}

//中间表
func TNBookCategory() string {
	return "md_book_category"
}

//book表
func TNBook() string {
	return "md_book"
}

//文档表
func TNDocuments() string {
	return "md_documents"
}

//文档存储表
func TNDocumentStore() string {
	return "md_document_store"
}

//附件表
func TNAttachment() string {
	return "md_attachment"
}

//关系表
func TNRelationship() string {
	return "md_relationship"
}

//成员表
func TNMembers() string {
	return "md_members"
}

//收藏表
func TNCollection() string {
	return "md_star"
}

//粉丝表
func TNFans() string {
	return "md_fans"
}

//评论表
func TNComments() string {
	return "md_comments"
}

//评分表
func TNScore() string {
	return "md_score"
}
