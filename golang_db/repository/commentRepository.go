package repository

import (
	"context"
	"golang_db/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Update(ctx context.Context, comment entity.Comment) (string, error)
	Delete(ctx context.Context, comment entity.Comment) (string, error)
	
	FindById(ctx context.Context, id int64) (entity.Comment, error)
	FindAll(ctx context.Context)([]entity.Comment, error)
	
}