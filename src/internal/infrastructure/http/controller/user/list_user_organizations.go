package user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) ListUserOrganizations(ctx echo.Context) error {
	c.log.Info(fmt.Sprintf("ListUserOrganizations invoked with id: %s", ctx.Param("id")))
	users, err := c.listUserOrganizationsUseCase.Execute(
		ctx.Param("id"),
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
