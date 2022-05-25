package service

// DyProjectCaseCreate 批量导入
// 导入kol自动填充部分信息
func (s *Service) DyProjectCaseCreate() error {
	// param := []*params.ArgDyProjectCase{}

	if err := s.dao.FamilyTree.Create(); err != nil {
		return err
	}

	return nil
}
