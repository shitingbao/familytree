package model

import "familytree/family/lib/interval"

type Member struct {
	ID           int               `form:"id" json:"id" gorm:"comumn:id"`
	Name         string            `form:"name" json:"name" gorm:"column:name"`
	ParentId     int               `form:"parentId" json:"parentId" gorm:"column:parent_id"`
	Path         string            `form:"path" json:"path" gorm:"column:path"`
	Sex          int               `form:"sex" json:"sex" gorm:"comumn:sex"`
	DateBirth    interval.TimeDate `form:"dateBirth" json:"dateBirth" gorm:"comumn:date_birth"`
	DateMarry    interval.TimeDate `form:"dateMarry" json:"dateMarry" gorm:"comumn:date_marry"`
	PlaceBirth   string            `form:"placeBirth" json:"placeBirth" gorm:"comumn:place_birth"`
	DateDeath    interval.TimeDate `form:"dateDeath" json:"dateDeath" gorm:"comumn:date_death"`
	PlaceDeath   string            `form:"placeDeath" json:"placeDeath" gorm:"comumn:place_death"`
	Content      string            `form:"content" json:"content" gorm:"comumn:content"`
	Honor        string            `form:"honor" json:"honor" gorm:"comumn:honor"`
	FamilySimple string            `form:"familySimple" json:"familySimple" gorm:"comumn:family_simple"`
	MarryId      int               `form:"marryId" json:"marryId" gorm:"comumn:marry_id"`
}

func (Member) TableName() string {
	return "member"
}
