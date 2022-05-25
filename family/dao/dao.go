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
	// postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
