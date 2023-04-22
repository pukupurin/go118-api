package main

import (
	infra "go-ent/infra/postgres"
	"go-ent/interface/router"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {

	db, err := infra.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// User関係のDI&ルーティングの初期化
	router.UserDIRouting(db, mux)

	// Start server
	http.ListenAndServe(
		":80",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
