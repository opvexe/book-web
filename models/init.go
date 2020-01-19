package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Category),
		new(Book),
		new(BookCategory),
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
