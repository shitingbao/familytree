package dao

import (
	"familytree/family/model"
	"familytree/family/params"

	"gorm.io/gorm"
)

type FamilyTreeCaseDao struct {
	*CurdDao
}

func NewFamilyTreeCaseDao(db *gorm.DB) *FamilyTreeCaseDao {
	return &FamilyTreeCaseDao{
		NewCurdDao(db, func() interface{} {
			return new(model.FamilyTreeBase)
		}),
	}
}

func (d *FamilyTreeCaseDao) Create(arg *model.FamilyTreeBase) error {
	return d.db.Table("").Create(arg).Error

}

func (d *FamilyTreeCaseDao) Delete(arg *params.ArgFamilyTree) error {
	return nil
}

func (d *FamilyTreeCaseDao) Update(id, status string) error {
	return nil
}

func (d *FamilyTreeCaseDao) BackupList() error {
	return nil
}
