package infra

import (
	"context"
	"go-ent/domain/model"
	"go-ent/domain/repository"
	"go-ent/ent"
)

type UserRepository struct {
	Conn *ent.Client
}

func NewUserRepository(conn *ent.Client) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Userの新規作成
func (ur *UserRepository) Create(ctx context.Context, user *model.User) (int, error) {
	u, err := ur.Conn.User.
		Create().
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}
