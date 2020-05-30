package sqlstore

import (
	"database/sql"
	"qask/internal/app/store"

	_ "github.com/go-sql-driver/mysql" //mysql
	_ "github.com/lib/pq"              //postgres
)

//Store is a struct with db
type Store struct {
	db                 *sql.DB
	questionRepository *QuestionRepository
	userRepository     *UserRepository
}

//New returns new struct Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Question returns question repository
func (s *Store) Question() store.QuestionRepository {
	if s.questionRepository != nil {
		return s.questionRepository
	}

	s.questionRepository = &QuestionRepository{
		store: s,
	}

	return s.questionRepository
}

//User returns user repository
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
