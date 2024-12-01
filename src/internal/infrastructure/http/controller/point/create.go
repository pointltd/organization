package point

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreatePointRequest struct {
	Name           string `json:"name" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
}

func (c *controller) CreatePoint(ctx echo.Context) error {
	request := new(CreatePointRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}

	organization, err := c.createPointUseCase.Execute(request.Name, request.OrganizationId)
	if err != nil {
		c.log.Error("Failed to create point for organization %s: %v", request.OrganizationId, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, organization)
}
