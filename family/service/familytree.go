package service

import (
	"familytree/family/model"
)

func (s *Service) FamilytreeCreate(arg *model.Member) error {
	return s.Familytree.Create(arg)
}

func (s *Service) FamilytreeUpdate(arg *model.Member) error {
	return s.Familytree.Update(arg)
}

func (s *Service) FamilytreeDelete(arg *model.Member) error {
	return s.Familytree.Delete(arg)
}

func (s *Service) FamilytreeList(arg *model.Member) ([]model.Member, error) {
	return s.Familytree.MemberList(arg)
}
