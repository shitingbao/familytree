package service

import "familytree/family/model"

func (s *Service) DyProjectCaseCreate() error {
	if err := s.dao.FamilyTree.Create(&model.FamilyTreeBase{}); err != nil {
		return err
	}

	return nil
}
