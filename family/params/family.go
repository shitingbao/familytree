package params

type ArgMember struct {
	ID         int    `form:"id" json:"id" gorm:"comumn:id"`
	Name       string `form:"name" json:"name" gorm:"column:name"`
	ParentId   int    `form:"parentId" json:"parentId" gorm:"column:parent_id"`
	Path       string `form:"path" json:"path" gorm:"column:path"`
	Sex        int    `form:"sex" json:"sex" gorm:"comumn:sex"`
	DateBirth  string `form:"dateBirth" json:"dateBirth" gorm:"comumn:date_birth"`
	DateMarry  string `form:"dateMarry" json:"dateMarry" gorm:"comumn:date_marry"`
	PlaceBirth string `form:"placeBirth" json:"placeBirth" gorm:"comumn:place_birth"`
	DateDeath  string `form:"dateDeath" json:"dateDeath" gorm:"comumn:date_death"`
	PlaceDeath string `form:"placeDeath" json:"placeDeath" gorm:"comumn:place_death"`
	Content    string `form:"content" json:"content" gorm:"comumn:content"`
	Honor      string `form:"honor" json:"honor" gorm:"comumn:honor"`
}

func (ArgMember) TableName() string {
	return "member"
}
