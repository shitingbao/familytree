package model

type Member struct {
	ID       string `json:"id"`
	ParentId string `json:"parent_id" gorm:"column:parent_id"`
	Name     string `json:"name" gorm:"column:name"`
}

func (Member) TableName() string {
	return "member"
}
