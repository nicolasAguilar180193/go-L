package player

func (s *Service) Delete(id string) (err error) {
	err = s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
