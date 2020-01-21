package models

import (
	"github.com/astaxie/beego/orm"
)

type AttachmentData struct {
	Attachment
	IsExist       bool
	BookName      string
	DocumentName  string
	FileShortSize string
	Account       string
	LocalHttpPath string
}

func NewAttachment() *Attachment {
	return &Attachment{}
}

func (m *Attachment) Insert() error {
	_, err := orm.NewOrm().Insert(m)
	return err
}

func (m *Attachment) Update() error {
	_, err := orm.NewOrm().Update(m)
	return err
}

func (m *Attachment) SelectByDocumentId(docId int) (attaches []*Attachment, err error) {
	_, err = orm.NewOrm().QueryTable(m.TableName()).Filter("document_id", docId).OrderBy("-attachment_id").All(&attaches)
	return
}
