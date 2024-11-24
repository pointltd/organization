package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/data"
	"net/http"
)

type CreateUserRequest struct {
	FirstName            string  `json:"first_name" validate:"required"`
	LastName             *string `json:"last_name"`
	Password             string  `json:"password" validate:"required"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqcsfield=Password"`
	Email                string  `json:"email" validate:"email"`
	Phone                *string `json:"phone" validate:"omitempty,e164"`
}

func (c *controller) CreateUser(ctx echo.Context) error {
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

	return ctx.JSON(http.StatusOK, user)
}
