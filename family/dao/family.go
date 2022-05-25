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

func (d *FamilyTreeCaseDao) Create() error {

	return nil
}

func (d *FamilyTreeCaseDao) Delete(arg *params.ArgFamilyTree) error {
	return nil
}

func (d *FamilyTreeCaseDao) UpdateBackupStatus(id, status string) error {
	return d.db.Table("dy_project_case").Where("id = ?", id).Update("backup_status", status).Error
}

func (d *FamilyTreeCaseDao) BackupList() error {

	return nil
}
