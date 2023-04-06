package main

import (
	infra "go-ent/infra/postgres"
	"go-ent/interface/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	db, err := infra.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// User関係のDI&ルーティングの初期化
	router.UserDIRouting(db, e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
