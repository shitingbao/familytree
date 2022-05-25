package dao

import (
	"familytree/family/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dao struct {
	c  *conf.Config
	db *gorm.DB

	FamilyTree *FamilyTreeCaseDao
}

func OpenDao(conf *conf.Config) (dao *gorm.DB, err error) {
	return gorm.Open(postgres.Open(conf.DB.Host), &gorm.Config{})
}

func NewDao(conf *conf.Config) *Dao {
	db, err := OpenDao(conf)
	if err != nil {
		panic(err)
	}

	return &Dao{
		c:          conf,
		db:         db,
		FamilyTree: NewFamilyTreeCaseDao(db),
	}
}
