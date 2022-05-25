package service

import (
	"familytree/family/conf"
	"familytree/family/dao"
)

type Service struct {
	c          *conf.Config
	dao        *dao.Data
	Familytree *dao.FamilyTreeCaseDao
}

func New(conf *conf.Config) *Service {
	d := dao.NewDao(conf)
	return &Service{
		c:          conf,
		dao:        d,
		Familytree: dao.NewFamilyTreeCaseDao(d),
	}
}
