package entity

type Answer struct {
	Id         int    `json:"id"`
	Answer     string `json:"answer"`
	Picture    string `json:"picture"`
	QuestionId int    `json:"questionId"`
}
