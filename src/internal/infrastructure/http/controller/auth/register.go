package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/data"
	"net/http"
	"os"
	"time"
)

type CreateUserRequest struct {
	FirstName            string  `json:"first_name" validate:"required"`
	LastName             *string `json:"last_name"`
	Password             string  `json:"password" validate:"required"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqcsfield=Password"`
	Email                string  `json:"email" validate:"email"`
	Phone                *string `json:"phone" validate:"omitempty,e164"`
}

func (c *controller) Register(ctx echo.Context) error {
	request := new(CreateUserRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}

	dto := data.CreateUserDTO{
		FirstName:            request.FirstName,
		LastName:             request.LastName,
		Password:             request.Password,
		PasswordConfirmation: request.PasswordConfirmation,
		Email:                request.Email,
	}

	user, err := c.createUserUseCase.Execute(dto)
	if err != nil {
		c.log.Error("Failed to create user: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
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
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user":  user,
	})
}
