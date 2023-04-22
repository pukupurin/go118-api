package router

import (
	"go-ent/ent"
	infra "go-ent/infra/postgres"
	"go-ent/interface/handler"
	"go-ent/usecase"
	"net/http"

	"github.com/bufbuild/connect-go"
)

func UserDIRouting(db *ent.Client, mux *http.ServeMux) {

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	i := connect.WithInterceptors()

	userPath, userServiceHandler := userHandler.UserServiceHandler(i)

	mux.Handle(userPath, userServiceHandler)
}
