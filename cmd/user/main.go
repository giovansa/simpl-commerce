package main

import (
	"database/sql"
	"fmt"
	"log"
	"simpl-commerce/handler"
	"simpl-commerce/internal"
	repository "simpl-commerce/repository/user"
	"simpl-commerce/transport/user"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	var server handler.IHandler = newServer(cfg)
	user.RegisterHandler(e, server)
	e.Logger.Fatal(e.Start(":" + cfg.App.Port))
}

func newServer(cfg internal.Config) *handler.Server {
	dbConn, err := sql.Open("postgres", fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DatabaseName))
	if err != nil {
		panic(err)
	}

	var repo repository.RepositoryInterface = repository.NewRepository(dbConn)
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(cfg, opts)
}
