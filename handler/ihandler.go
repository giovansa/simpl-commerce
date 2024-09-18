package handler

import "github.com/labstack/echo/v4"

type IHandler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	GetProfile(ctx echo.Context) error
}
