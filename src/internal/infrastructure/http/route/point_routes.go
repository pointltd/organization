package route

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
)

func RegisterPointRoutes(group *echo.Group, controller controller.PointController, jwtMiddleware echo.MiddlewareFunc) {
	pointGroup := group.Group("/points")
	//pointGroup.Use(jwtMiddleware)
	pointGroup.POST("", controller.CreatePoint)
}
