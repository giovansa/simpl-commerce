package handler

import (
	"net/http"
	"simpl-commerce/internal"
	"simpl-commerce/model/common"
	"simpl-commerce/model/user"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *Server) Register(ctx echo.Context) error {
	user := new(user.RegisterUserRequest)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := user.Validate()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userInput, err := user.ToDAO()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userID, err := s.Repository.RegisterUser(ctx.Request().Context(), userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": map[string]string{
		"user_id": userID,
	}})
}

func (s *Server) Login(ctx echo.Context) error {
	req := new(user.UserLoginRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	ctx2 := ctx.Request().Context()
	userDAO, err := s.Repository.GetUserByPhone(ctx2, strings.TrimSpace(req.Phone))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	user := user.FromRepoUser(userDAO)
	err = user.CheckLogin(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	tokenString, err := internal.GenerateJWTToken(user, s.Cfg.App.Env)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func (s *Server) GetProfile(ctx echo.Context) error {
	claimUser := ctx.Get("claims").(*common.Claims)
	if claimUser == nil {
		return ctx.JSON(http.StatusForbidden, map[string]struct{}{})
	}

	userDAO, err := s.Repository.GetUserByPhone(ctx.Request().Context(), claimUser.Phone)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	user := user.FromRepoUser(userDAO)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": user.ToProfileResp(),
	})
}
