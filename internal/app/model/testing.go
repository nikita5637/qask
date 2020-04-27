package model

//TestMathProblem is a function that return test math problem
func TestMathProblem() *Question {
	m := &Question{
		Sitename: "Тестовая база",
		Question: "Тестовая задача",
		Answer:   "Тестовый ответ к задаче",
		Comment:  "Тестовый комментарий к задаче",
	}

	return m
}

//TestQuestion is a function that returns test question
func TestQuestion() *Question {
	q := &Question{
		Sitename: "Тестовая база",
		Question: "Тестовый вопрос",
		Answer:   "Тестовый ответ",
		Comment:  "Тестовый комментарий",
	}

	return q
}

//TestUser is a function that returns test user
func TestUser() *User {
	u := &User{
		userPublic{
			FirstName: "TestUser_FirstName",
		},
		userPrivate{
			UserName: "TestUser_UserName",
			TgID:     100,
		},
	}

	return u
}
