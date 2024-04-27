package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
}

func NewClaims(claimMap map[string]interface{}) (*Claims, error) {
	ExpiresAt, ok := claimMap["exp"].(int64)
	if !ok {
		return nil, jwt.ValidationError{}
	}

	Issuer, ok := claimMap["iss"].(string)
	if !ok {
		return nil, jwt.ValidationError{}
	}

	Subject, ok := claimMap["sub"].(string)
	if !ok {
		return nil, jwt.ValidationError{}
	}

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
			Issuer:    Issuer,
			Subject:   Subject,
		},
	}

	return claims, nil
}
