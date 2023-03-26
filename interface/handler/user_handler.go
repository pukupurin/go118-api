package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go-ent/usecase"

	"github.com/labstack/echo/v4"
)

// UserHander User handerのinterface
type UserHandler interface {
	CreateUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	GetUserList() echo.HandlerFunc
	GetUserOne() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler User handler のコンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// req & res json -------------
type DefaultResponse struct {
	Message string `json:"message"`
}

// ----------------------------

func (h *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res DefaultResponse

		postdata := new(usecase.ReqCreateUpdateUser)
		err := c.Bind(postdata)
		if err != nil {
			res.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, res)
		}

		err = h.userUsecase.CreateUser(c.Request().Context(), *postdata)
		if err != nil {
			res.Message = "create user error"
			fmt.Println(err.Error())
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Message = "success"

		return c.JSON(http.StatusOK, res)
	}
}

func (h *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res DefaultResponse

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			res.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, res)
		}

		postdata := new(usecase.ReqCreateUpdateUser)
		if err := c.Bind(postdata); err != nil {
			res.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, res)
		}

		err = h.userUsecase.UpdateUser(c.Request().Context(), id, *postdata)
		if err != nil {
			res.Message = "update user error"
			fmt.Println(err.Error())
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Message = "success"

		return c.JSON(http.StatusOK, res)
	}
}

func (h *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res DefaultResponse

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			res.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, res)
		}

		err = h.userUsecase.DeleteUser(c.Request().Context(), id)
		if err != nil {
			res.Message = "delete user error"
			fmt.Println(err.Error())
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Message = "success"

		return c.JSON(http.StatusOK, res)
	}
}

func (h *userHandler) GetUserList() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res usecase.ResGetUserList
		var resError DefaultResponse

		postdata := new(usecase.ReqGetUserList)
		err := c.Bind(postdata)
		if err != nil {
			resError.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, resError)
		}

		res, err = h.userUsecase.GetUserList(c.Request().Context(), postdata.Limit, postdata.Offset)
		if err != nil {
			resError.Message = "get user error"
			fmt.Println(err.Error())
			return c.JSON(http.StatusBadRequest, resError)
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (h *userHandler) GetUserOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res usecase.ResGetUser
		var resError DefaultResponse

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			resError.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, resError)
		}

		res, err = h.userUsecase.GetUserOne(c.Request().Context(), id)
		if err != nil {
			resError.Message = "get user error"
			fmt.Println(err.Error())
			return c.JSON(http.StatusBadRequest, resError)
		}

		return c.JSON(http.StatusOK, res)
	}
}
