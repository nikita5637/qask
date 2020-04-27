package www

//MathProblemsRepository implemets getting math problem from internet sites
type MathProblemsRepository struct {
	db map[int]SiteRepository
}

//GetMathProblem ...
func (m *MathProblemsRepository) GetMathProblem() (interface{}, error) {
	return m.db[0].GetQuestion()
}
