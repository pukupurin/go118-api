package infra

import (
	"context"
	"database/sql"
	"fmt"
	"go-ent/ent"
	"os"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func OpenDB() (*ent.Client, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dbConf := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pass)
	fmt.Println(dbConf)

	db, err := sql.Open("postgres", dbConf)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Driver(drv))
	entOptions = append(entOptions, ent.Debug())

	return ent.NewClient(entOptions...), nil
}
