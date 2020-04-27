package questions

//QuestionsRepository implements getting a question
type QuestionsRepository interface {
	GetQuestion() (interface{}, error)
}

//MathProblemsRepository implements getting a math problem
type MathProblemsRepository interface {
	GetMathProblem() (interface{}, error)
}
