package teststore

import "qask/internal/app/model"

//QuestionRepository is a questions store for testing
type QuestionRepository struct {
	questions map[int]*model.Question
}
