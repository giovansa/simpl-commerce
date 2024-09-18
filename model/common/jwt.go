package common

import "github.com/golang-jwt/jwt"

type Claims struct {
	StandardClaims jwt.StandardClaims
	Phone          string
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}
