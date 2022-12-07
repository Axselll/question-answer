package dtos

type CreateQuestionRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
