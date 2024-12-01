package organization

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) ListOrganizationPoints(ctx echo.Context) error {
	users, err := c.listOrganizationPointsUseCase.Execute(
		ctx.Param("id"),
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
