package www

import (
	"qask/internal/app/questions"
	"qask/internal/app/questions/www/db_chgk_info"
	"qask/internal/app/questions/www/problems_ru"
)

//Questions ...
type Questions struct {
	questionsRepository    *QuestionsRepository
	mathProblemsRepository *MathProblemsRepository
}

//New ...
func New() *Questions {
	return &Questions{}
}

//Questions ...
func (q *Questions) Questions() questions.QuestionsRepository {
	if q.questionsRepository != nil {
		return q.questionsRepository
	}

	q.questionsRepository = &QuestionsRepository{}

	q.questionsRepository.db = make(map[int]SiteRepository)

	q.questionsRepository.db[0] = db_chgk_info.New()

	return q.questionsRepository
}

//MathProblems ...
func (q *Questions) MathProblems() questions.MathProblemsRepository {
	if q.mathProblemsRepository != nil {
		return q.mathProblemsRepository
	}

	q.mathProblemsRepository = &MathProblemsRepository{}

	q.mathProblemsRepository.db = make(map[int]SiteRepository)

	q.mathProblemsRepository.db[0] = problems_ru.New()

	return q.mathProblemsRepository
}
