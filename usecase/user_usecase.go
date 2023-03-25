package usecase

import (
	"context"
	"go-ent/domain/model"
	"go-ent/domain/repository"
)

// UserUsecase User 関係のusecaseのinterface
type UserUsecase interface {
	CreateUser(ctx context.Context, reqData ReqCreateUser) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase User usecaseのコンストラクタ
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

type ReqCreateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUser User を追加する
func (u *userUsecase) CreateUser(ctx context.Context, reqData ReqCreateUser) error {

	user, err := model.NewUser(reqData.Name, reqData.Age)
	if err != nil {
		return err
	}

	_, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
