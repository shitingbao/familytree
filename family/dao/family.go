package dao

import (
	"familytree/family/model"
	"familytree/family/params"

	"github.com/pborman/uuid"
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
	arg.ID = uuid.New()
	return d.Data.db.Table("member").Create(arg).Error
}

///
func (d *FamilyTreeCaseDao) Delete(arg *params.ArgMemeber) error {
	return d.Data.DB().Table("member").Where("id", arg.ID).Update("is_deleted", isDeletedYes).Error
}

func (d *FamilyTreeCaseDao) Update(arg *params.ArgMemeber) error {
	return d.Data.DB().Table("member").Where("id", arg.ID).Updates(arg).Error
}

func (d *FamilyTreeCaseDao) MemberList(arg *params.ArgMemeber) ([]model.Member, error) {
	list := []model.Member{}
	db := d.Data.DB().Table("member")
	if arg.ID == "" {
		db = db.Where("parent_id is null")
	} else {
		db = db.Where("parent_id", arg.ID)
	}
	err := db.Scan(&list).Error
	return list, err
}
