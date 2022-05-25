package service

import (
	"familytree/family/model"
	"familytree/family/params"
)

func (s *Service) FamilytreeCreate(arg *model.Member) error {
	if err := s.Familytree.Create(arg); err != nil {
		return err
	}
	return nil
}

func (s *Service) FamilytreeUpdate(arg *params.ArgMemeber) error {
	if err := s.Familytree.Update(arg); err != nil {
		return err
	}
	return nil
}

func (s *Service) FamilytreeDelete(arg *params.ArgMemeber) error {
	if err := s.Familytree.Delete(arg); err != nil {
		return err
	}
	return nil
}
