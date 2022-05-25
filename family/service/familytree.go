package service

import (
	"familytree/family/model"
	"familytree/family/params"
)

func (s *Service) FamilytreeCreate(arg *model.Member) error {
	return s.Familytree.Create(arg)
}

func (s *Service) FamilytreeUpdate(arg *params.ArgMemeber) error {
	return s.Familytree.Update(arg)
}

func (s *Service) FamilytreeDelete(arg *params.ArgMemeber) error {
	return s.Familytree.Delete(arg)
}

func (s *Service) FamilytreeList(arg *params.ArgMemeber) ([]model.Member, error) {
	return s.Familytree.MemberList(arg)
}
