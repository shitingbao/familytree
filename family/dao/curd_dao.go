package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CurdDao struct {
	modelFunc func() interface{}
	db        *gorm.DB
}

func NewCurdDao(db *gorm.DB, f func() interface{}) *CurdDao {
	return &CurdDao{
		db:        db,
		modelFunc: f,
	}
}
