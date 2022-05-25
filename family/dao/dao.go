package dao

import (
	"fmt"

	"gorm.io/gorm/logger"

	"familytree/family/conf"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	c  *conf.Config
	db *gorm.DB

	FamilyTree *FamilyTreeCaseDao
}

func OpenDao(conf *conf.Config) (dao *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Database)
	loggerConfig := logger.Default.LogMode(logger.Silent)
	if conf.DB.LogMode > 0 {
		loggerConfig = logger.Default.LogMode(logger.LogLevel(conf.DB.LogMode))
	}
	dao, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConfig,
	})
	if err != nil {
		return nil, err
	}
	db, err := dao.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	return
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
