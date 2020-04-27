package questions

//Questions is a interface
type Questions interface {
	Questions() QuestionsRepository
	MathProblems() MathProblemsRepository
}
