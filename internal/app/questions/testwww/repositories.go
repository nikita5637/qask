package testwww

import (
	"qask/internal/app/model"
)

//QuestionsRepository implements getting a questions from test store
type QuestionsRepository struct {
}

//MathProblemsRepository implements getting a math problem from test store
type MathProblemsRepository struct {
}

//GetQuestion is a function, that returns test question
func (q *QuestionsRepository) GetQuestion() (interface{}, error) {
	testQuestion := model.TestQuestion()

	return testQuestion, nil
}

//GetMathProblem is a function, that returns returns test math problem
func (m *MathProblemsRepository) GetMathProblem() (interface{}, error) {
	testMathProblem := model.TestMathProblem()

	return testMathProblem, nil
}
