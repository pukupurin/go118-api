package router

import (
	"go-ent/ent"
	infra "go-ent/infra/postgres"
	"go-ent/interface/handler"
	"go-ent/usecase"

	"github.com/labstack/echo/v4"
)

func UserDIRouting(db *ent.Client, e *echo.Echo) {

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e.POST("/users", userHandler.CreateUser())
	e.PUT("/users/:id", userHandler.UpdateUser())
	e.DELETE("/users/:id", userHandler.DeleteUser())
	e.GET("/users", userHandler.GetUserList())
	e.GET("/users/:id", userHandler.GetUserOne())
}
