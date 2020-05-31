package teststore

import (
	"qask/internal/app/model"
	"qask/internal/app/store"
)

//Store ...
type Store struct {
	userRepository     *UserRepository
	questionRepository *QuestionRepository
	reportRepository   *ReportRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		users: make(map[int64]*model.User),
	}

	return s.userRepository
}

//Question ...
func (s *Store) Question() store.QuestionRepository {
	if s.questionRepository != nil {
		return s.questionRepository
	}

	s.questionRepository = &QuestionRepository{
		questions: make(map[int64]*model.Question),
	}

	return s.questionRepository
}

//Report ...
func (s *Store) Report() store.ReportRepository {
	if s.reportRepository != nil {
		return s.reportRepository
	}

	s.reportRepository = &ReportRepository{}

	return s.reportRepository
}
