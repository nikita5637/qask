package testwww

import "qask/internal/app/questions"

//Questions is a struct with repositories
type Questions struct {
	mathProblemsRepository *MathProblemsRepository
	questionsRepository    *QuestionsRepository
}

//New function create new testwww repository
func New() *Questions {
	return &Questions{}
}

//MathProblems is a function that return MathProblemsRepository
func (q *Questions) MathProblems() questions.MathProblemsRepository {
	if q.mathProblemsRepository != nil {
		return q.mathProblemsRepository
	}

	q.mathProblemsRepository = &MathProblemsRepository{}

	return q.mathProblemsRepository
}

//Questions is a function that returns QuiestionsRepository
func (q *Questions) Questions() questions.QuestionsRepository {
	if q.questionsRepository != nil {
		return q.questionsRepository
	}

	q.questionsRepository = &QuestionsRepository{}

	return q.questionsRepository
}
