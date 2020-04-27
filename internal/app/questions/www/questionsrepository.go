package www

//QuestionsRepository ...
type QuestionsRepository struct {
	db map[int]SiteRepository
}

//GetQuestion ...
func (q *QuestionsRepository) GetQuestion() (interface{}, error) {
	return q.db[0].GetQuestion()
}
