package model

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name" gorm:"column:name"`
}

func (Member) TableName() string {
	return "member"
}
