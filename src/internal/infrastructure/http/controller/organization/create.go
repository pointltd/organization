package organization

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/http"
	nethttp "net/http"
)

type CreateOrganizationRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *controller) CreateOrganization(ctx echo.Context) error {
	request := new(CreateOrganizationRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}

	claims := http.GetClaims(ctx)
	ownerID := claims.ID

	organization, err := c.createOrganizationUseCase.Execute(request.Name, ownerID)
	if err != nil {
		c.log.Error("Failed to create organization for user %s: %v", ownerID, err)
		return ctx.JSON(nethttp.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(nethttp.StatusOK, organization)
}
