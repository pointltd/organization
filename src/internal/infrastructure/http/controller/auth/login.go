package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *controller) Login(ctx echo.Context) error {
	var req AuthenticateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	user, err := c.authenticateUserUseCase.Execute(req.Username, req.Password)

	if err != nil || user == nil {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		user.Id,
		user.Contacts.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(c.config.JwtSecret()))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user":  user,
	})
}
