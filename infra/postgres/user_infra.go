package infra

import (
	"context"
	"go-ent/domain/model"
	"go-ent/domain/repository"
	"go-ent/ent"
	entuser "go-ent/ent/user"
	"time"

	"github.com/samber/lo"
)

type UserRepository struct {
	Conn *ent.Client
}

func NewUserRepository(conn *ent.Client) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Userの新規作成
func (r *UserRepository) Create(ctx context.Context, data *model.User) (int, error) {
	u, err := r.Conn.User.
		Create().
		SetName(data.Name).
		SetAge(data.Age).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Update Userの更新
func (r *UserRepository) Update(ctx context.Context, data *model.User) error {
	_, err := r.Conn.User.
		UpdateOneID(data.ID).
		SetName(data.Name).
		SetAge(data.Age).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Delete Userの削除
func (r *UserRepository) Delete(ctx context.Context, userID int) error {
	_, err := r.Conn.User.
		UpdateOneID(userID).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetList Userの複数件取得
func (r *UserRepository) GetList(ctx context.Context, limit int, offset int) ([]model.User, error) {
	resUsers := make([]model.User, 0, limit)

	us, err := r.Conn.User.
		Query().
		Where(entuser.DeletedAtIsNil()).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return resUsers, err
	}

	resUsers = lo.Map(us, func(u *ent.User, _ int) model.User {
		return model.User{
			ID:   u.ID,
			Name: u.Name,
			Age:  u.Age,
		}
	})

	return resUsers, nil
}

// GetOne Userを1件取得
func (r *UserRepository) GetOne(ctx context.Context, userID int) (model.User, error) {
	var resUser model.User

	u, err := r.Conn.User.
		Query().
		Where(entuser.DeletedAtIsNil()).
		Where(entuser.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		return resUser, err
	}

	resUser = model.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}

	return resUser, nil
}
