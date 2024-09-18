package internal

import (
	"crypto/rsa"
	"os"
	"simpl-commerce/model/common"
	"simpl-commerce/model/user"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(user user.User, env string) (string, error) {
	privateKey, err := loadPrivateKey(env)
	if err != nil {
		return "", err
	}

	claims := &common.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		Phone: user.Phone,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func loadPrivateKey(env string) (*rsa.PrivateKey, error) {
	if env == "test" {
		keyDer, err := os.ReadFile("../private.pem")
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPrivateKeyFromPEM(keyDer)
	}
	keyDer, err := os.ReadFile("./private.pem")
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyDer)
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	keyDer, err := os.ReadFile("./public.pem")
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyDer)
}
