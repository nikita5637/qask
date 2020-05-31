package store

//Store is a interface for a interaction with databases
type Store interface {
	User() UserRepository
	Question() QuestionRepository
	Report() ReportRepository
}
