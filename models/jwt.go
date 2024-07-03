package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	SurName         string `json:"surname"`
	Role            string `json:"role"`
	CompanyID       string `json:"companyID"`
	CompanyCategory string `json:"companyCategory"`
	jwt.RegisteredClaims
}

type Token struct {
	Token string `json:"token"`
}
