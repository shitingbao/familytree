package dao

import (
	"familytree/family/model"
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

///
func (d *FamilyTreeCaseDao) Delete(arg *model.Member) error {
	return d.Data.DB().Table("member").Where("id", arg.ID).Update("is_deleted", isDeletedYes).Error
}

func (d *FamilyTreeCaseDao) Update(arg *model.Member) error {
	return d.Data.DB().Table("member").Where("id", arg.ID).Updates(arg).Error
}

func (d *FamilyTreeCaseDao) MemberList(arg *model.Member) ([]model.Member, error) {
	list := []model.Member{}
	db := d.Data.DB().Table("member").Where("parent_id", arg.ID)

	err := db.Scan(&list).Error
	return list, err
}
