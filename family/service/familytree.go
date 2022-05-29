package service

import (
	"familytree/family/model"
	"familytree/family/params"
)

// FamilytreeCreate 根据父节点构建一个新的子节点
func (s *Service) FamilytreeCreate(arg *model.Member) error {
	return s.Familytree.Create(arg)
}

func (s *Service) FamilytreeUpdate(arg *params.ArgMember) error {
	return s.Familytree.Update(arg)
}

func (s *Service) FamilytreeDelete(arg *model.Member) error {
	return s.Familytree.Delete(arg)
}

func (s *Service) FamilytreeList(arg *model.Member) ([]model.Member, error) {
	return s.Familytree.MemberList(arg)
}

func (s *Service) FamilytreeLast(arg *model.Member) ([]model.Member, error) {
	return s.Familytree.MemberLast(arg)
}

func (s *Service) SearchMember(arg *model.Member) ([]model.Member, error) {
	return s.Familytree.SearchMember(arg)
}
