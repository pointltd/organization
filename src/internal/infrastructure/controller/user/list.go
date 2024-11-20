package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) ListUsers(ctx echo.Context) error {
	users, err := c.listUsersUseCase.Execute()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
