package usecase

import (
	"context"
	"qa/src/dtos"
	"qa/src/repository"
)

type QuestionUsecase interface {
	List(ctx context.Context, request dtos.CreateQuestionRequest) dtos.QuestionResponse
	Create(ctx context.Context) []dtos.QuestionResponse
}

type QuestionUsecaseImpl struct {
	QuestionRepository repository.QuestionRepository
}

func (questionUsecase *QuestionUsecaseImpl) List(ctx context.Context, request dtos.CreateQuestionRequest) dtos.QuestionResponse {
	panic("not implemented") // TODO: Implement
}

func (questionUsecase *QuestionUsecaseImpl) Create(ctx context.Context) []dtos.QuestionResponse {
	panic("not implemented") // TODO: Implement
}
