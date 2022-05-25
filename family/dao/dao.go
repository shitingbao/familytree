package dao

import (
	"familytree/family/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	c  *conf.Config
	db *gorm.DB
}

func OpenDao(conf *conf.Config) (dao *gorm.DB, err error) {
	return gorm.Open(postgres.Open(conf.DB.Host), &gorm.Config{})
}

func NewDao(conf *conf.Config) *Data {
	db, err := OpenDao(conf)
	if err != nil {
		panic(err)
	}

	return &Data{
		c:  conf,
		db: db,
	}
}

// 默认排除删除的数据
func (d *Data) DB() *gorm.DB {
	return d.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_deleted=?", 0)
	})
}
