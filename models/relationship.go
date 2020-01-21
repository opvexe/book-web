package models

import (
	"github.com/astaxie/beego/orm"
)

func NewRelationship() *Relationship {
	return &Relationship{}
}

func (m *Relationship) Select(bookId, memberId int) (*Relationship, error) {
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("book_id", bookId).Filter("member_id", memberId).One(m)
	return m, err
}

func (m *Relationship) SelectRoleId(bookId, memberId int) (int, error) {
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("book_id", bookId).Filter("member_id", memberId).One(m, "role_id")
	if err != nil {
		return 0, err
	}
	return m.RoleId, nil
}

func (m *Relationship) Insert() error {
	_, err := orm.NewOrm().Insert(m)
	return err
}

func (m *Relationship) Update() error {
	_, err := orm.NewOrm().Update(m)
	return err
}
