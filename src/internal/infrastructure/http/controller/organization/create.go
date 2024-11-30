package organization

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	http_pkg "github.com/pointltd/organization/internal/infrastructure/http"
	"net/http"
)

type CreateOrganizationRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *controller) CreateOrganization(ctx echo.Context) error {
	request := new(CreateOrganizationRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*http_pkg.JwtCustomClaims)
	ownerID := claims.ID

	organization, err := c.createOrganizationUseCase.Execute(request.Name, ownerID)
	if err != nil {
		c.log.Error("Failed to create organization for user %s: %v", ownerID, err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, organization)
}
