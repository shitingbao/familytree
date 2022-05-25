package service

import (
	"familytree/family/conf"
	"familytree/family/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

func New(conf *conf.Config) *Service {
	return &Service{
		c:   conf,
		dao: dao.NewDao(conf),
	}
}
