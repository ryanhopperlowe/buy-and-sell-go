package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func DefaultClaims(subject string) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "buy-and-sell",
			Subject:   subject,
			NotBefore: time.Now().Unix(),
		},
	}
}
