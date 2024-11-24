package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateUserRequest struct {
	FirstName            string  `json:"first_name" validate:"required"`
	LastName             *string `json:"last_name"`
	Password             string  `json:"password" validate:"required"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqcsfield=Password"`
	Email                string  `json:"email" validate:"email"`
}

func (c *controller) CreateUser(ctx echo.Context) error {
	request := new(CreateUserRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, "")
}
