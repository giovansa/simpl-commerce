package user

import (
	"simpl-commerce/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Handler handler.IHandler
	Router  *echo.Echo
}

func RegisterHandler(e *echo.Echo, handler handler.IHandler) {
	e.POST("/user", handler.Register)
	e.POST("/login", handler.Login)

	// e.POST("/order", handler.Login, transport.AuthMiddleware)
}
