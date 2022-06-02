package model

import "familytree/family/lib/interval"

type Member struct {
	ID           int               `form:"id" json:"id" gorm:"column:id"`
	Name         string            `form:"name" json:"name" gorm:"column:name"`
	ParentId     int               `form:"parentId" json:"parentId" gorm:"column:parent_id"`
	Path         string            `form:"path" json:"path" gorm:"column:path"`
	Sex          string            `form:"sex" json:"sex" gorm:"column:sex"`
	DateBirth    interval.TimeDate `form:"dateBirth" json:"dateBirth" gorm:"column:date_birth"`
	DateMarry    interval.TimeDate `form:"dateMarry" json:"dateMarry" gorm:"column:date_marry"`
	PlaceBirth   string            `form:"placeBirth" json:"placeBirth" gorm:"column:place_birth"`
	DateDeath    interval.TimeDate `form:"dateDeath" json:"dateDeath" gorm:"column:date_death"`
	PlaceDeath   string            `form:"placeDeath" json:"placeDeath" gorm:"column:place_death"`
	Content      string            `form:"content" json:"content" gorm:"column:content"`
	Honor        string            `form:"honor" json:"honor" gorm:"column:honor"`
	FamilySimple string            `form:"familySimple" json:"familySimple" gorm:"column:family_simple"`
	MarryId      int               `form:"marryId" json:"marryId" gorm:"column:marry_id"`
}

func (Member) TableName() string {
	return "member"
}
