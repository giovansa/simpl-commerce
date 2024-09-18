package transport

import (
	"errors"
	"net/http"
	"simpl-commerce/internal"
	"simpl-commerce/model/common"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusForbidden, map[string]struct{}{})
		}
		publicKey, err := internal.LoadPublicKey()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]struct{}{})
		}
		token, err := jwt.ParseWithClaims(tokenString, &common.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return c.JSON(http.StatusForbidden, map[string]struct{}{})
			}
			return c.JSON(http.StatusForbidden, map[string]struct{}{})
		}
		if !token.Valid {
			return c.JSON(http.StatusForbidden, map[string]struct{}{})
		}

		claims := token.Claims.(*common.Claims)
		c.Set("claims", claims)

		return next(c)
	}
}
