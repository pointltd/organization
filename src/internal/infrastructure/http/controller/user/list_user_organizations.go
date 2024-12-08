package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) ListUserOrganizations(ctx echo.Context) error {
	c.log.Debug("ListUserOrganizations invoked")
	users, err := c.listUserOrganizationsUseCase.Execute(
		ctx.Param("id"),
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
