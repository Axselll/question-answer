package repository

import (
	"context"
	"database/sql"
	"qa/src/entity"
)

type QuestionRepository interface {
	List(ctx context.Context, tx *sql.Tx) []entity.Question
	Create(ctx context.Context, tx *sql.Tx, question entity.Question) entity.Question
}

type QuestionRepositoryImpl struct {
}

func (questionRepository *QuestionRepositoryImpl) List(ctx context.Context, tx *sql.Tx) []entity.Question {
	panic("not implemented") // TODO: Implement
}

func (questionRepository *QuestionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, question entity.Question) entity.Question {
	panic("not implemented") // TODO: Implement
}
