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

func (d *FamilyTreeCaseDao) BackupList() error {
	return nil
}
