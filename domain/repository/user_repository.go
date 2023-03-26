package repository

import (
	"context"
	"go-ent/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, data *model.User) (int, error)
	Update(ctx context.Context, data *model.User) error
	Delete(ctx context.Context, userID int) error
	GetList(ctx context.Context, limit int, offset int) ([]model.User, error)
	GetOne(ctx context.Context, userID int) (model.User, error)
}
