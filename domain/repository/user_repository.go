package repository

import (
	"context"
	"go-ent/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int, error)
}
