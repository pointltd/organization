package http

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GetJwtConfig(jwtSecret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(jwtSecret),
	}
}

func GetClaims(ctx echo.Context) JwtCustomClaims {
	user := ctx.Get("user").(*jwt.Token)
	return *user.Claims.(*JwtCustomClaims)
}
