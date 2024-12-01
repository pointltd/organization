package route

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
)

func RegisterOrganizationRoutes(
	group *echo.Group,
	controller controller.OrganizationController,
	jwtMiddleware echo.MiddlewareFunc,
) {
	organizationGroup := group.Group("/organizations")
	organizationGroup.Use(jwtMiddleware)
	organizationGroup.POST("", controller.CreateOrganization)
	organizationGroup.GET("/:id/points", controller.ListOrganizationPoints)
}
