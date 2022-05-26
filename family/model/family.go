package model

import "familytree/family/lib/interval"

type Member struct {
	ID         int               `json:"id" gorm:"comumn:id"`
	Name       string            `json:"name" gorm:"column:name"`
	ParentId   int               `json:"parentId" gorm:"column:parent_id"`
	Path       string            `json:"path" gorm:"column:path"`
	Sex        int               `json:"sex" gorm:"comumn:sex"`
	DateBirth  interval.TimeDate `json:"dateBirth" gorm:"comumn:date_birth"`
	DateMarry  interval.TimeDate `json:"dateMarry" gorm:"comumn:date_marry"`
	PlaceBirth string            `json:"placeBirth" gorm:"comumn:place_birth"`
	DateDeath  interval.TimeDate `json:"dateDeath" gorm:"comumn:date_death"`
	PlaceDeath string            `json:"placeDeath" gorm:"comumn:place_death"`
	Content    string            `json:"content" gorm:"comumn:content"`
	Honor      string            `json:"honor" gorm:"comumn:honor"`
}

func (Member) TableName() string {
	return "member"
}
