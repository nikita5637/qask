package teststore

import (
	"qask/internal/app/model"
	"qask/internal/app/store"
)

type Store struct {
	userRepository     *UserRepository
	questionRepository *QuestionRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Question() store.QuestionRepository {
	if s.questionRepository != nil {
		return s.questionRepository
	}

	s.questionRepository = &QuestionRepository{
		questions: make(map[int]*model.Question),
	}

	return s.questionRepository
}
