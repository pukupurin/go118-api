package handler

import (
	"context"
	"fmt"
	"net/http"

	userv1 "go-ent/gen/user/v1"
	"go-ent/gen/user/v1/userv1connect"
	"go-ent/usecase"

	"github.com/bufbuild/connect-go"
	"github.com/samber/lo"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler User handler のコンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) UserServiceHandler(i connect.HandlerOption) (string, http.Handler) {
	return userv1connect.NewUserServiceHandler(h, i)
}

func (h *UserHandler) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[emptypb.Empty], error) {
	reqData := usecase.ReqCreateUpdateUser{
		Name: req.Msg.Name,
		Age:  int(req.Msg.Age),
	}

	err := h.userUsecase.CreateUser(ctx, reqData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[emptypb.Empty], error) {
	reqData := usecase.ReqCreateUpdateUser{
		Name: req.Msg.Name,
		Age:  int(req.Msg.Age),
	}

	err := h.userUsecase.UpdateUser(ctx, int(req.Msg.Id), reqData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {

	err := h.userUsecase.DeleteUser(ctx, int(req.Msg.Id))
	if err != nil {
		fmt.Println(err.Error())
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (h *UserHandler) GetUserList(ctx context.Context, req *connect.Request[userv1.GetUserListRequest]) (*connect.Response[userv1.GetUserListResponse], error) {

	resUsers, err := h.userUsecase.GetUserList(ctx, int(req.Msg.Limit), int(req.Msg.Offset))
	if err != nil {
		fmt.Println(err.Error())
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&userv1.GetUserListResponse{
		Users: lo.Map(resUsers.Users, func(u usecase.ResGetUser, _ int) *userv1.GetUserOneResponse {
			return &userv1.GetUserOneResponse{
				Id:   int32(u.ID),
				Name: u.Name,
				Age:  int32(u.Age),
			}
		}),
	}), nil
}

func (h *UserHandler) GetUserOne(ctx context.Context, req *connect.Request[userv1.GetUserOneRequest]) (*connect.Response[userv1.GetUserOneResponse], error) {

	resUser, err := h.userUsecase.GetUserOne(ctx, int(req.Msg.Id))
	if err != nil {
		fmt.Println(err.Error())
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&userv1.GetUserOneResponse{
		Id:   int32(resUser.ID),
		Name: resUser.Name,
		Age:  int32(resUser.Age),
	}), nil
}
