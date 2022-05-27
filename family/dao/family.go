package dao

import (
	"familytree/family/model"
	"familytree/family/params"
	"strconv"
)

var (
	isDeletedYes = 1
)

type FamilyTreeCaseDao struct {
	Data *Data
}

func NewFamilyTreeCaseDao(da *Data) *FamilyTreeCaseDao {
	return &FamilyTreeCaseDao{
		Data: da,
	}
}

func (d *FamilyTreeCaseDao) Create(arg *model.Member) error {
	return d.Data.db.Table("member").Create(arg).Error
}

// Delete 删除注意，删除本节点，以及对应后续子节点
func (d *FamilyTreeCaseDao) Delete(arg *model.Member) error {
	return d.Data.DB().Table("member").Where("id = ? or path = ? or path like ?", arg.ID, strconv.Itoa(arg.ID), strconv.Itoa(arg.ID)+",%").Update("is_deleted", isDeletedYes).Error
}

func (d *FamilyTreeCaseDao) Update(arg *params.ArgMember) error {
	return d.Data.DB().Debug().Table("member").Where("id", arg.ID).Updates(arg).Error
}

func (d *FamilyTreeCaseDao) MemberList(arg *model.Member) ([]model.Member, error) {
	list := []model.Member{}
	db := d.Data.DB().Table("member").Where("parent_id", arg.ID)

	err := db.Scan(&list).Error
	return list, err
}

// MemberLast 找到一棵完整的树
// 注意节点前缀 比如 6和60
func (d *FamilyTreeCaseDao) MemberLast(arg *model.Member) ([]model.Member, error) {
	list := []model.Member{}
	id := arg.ID
	if arg.ID == 0 {
		d.Data.DB().Select("id").Table("member").Where("parent_id = 0").Find(&id)
	}
	db := d.Data.DB().Debug().Table("member").
		Where("path like ? or id = ? or path = ?", strconv.Itoa(id)+",%", id, strconv.Itoa(id))
	err := db.Scan(&list).Error
	return list, err
}
