package route

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
)

func RegisterUserRoutes(group *echo.Group, controller controller.UserController, jwtMiddleware echo.MiddlewareFunc) {
	userGroup := group.Group("/users")
	//userGroup.Use(jwtMiddleware)
	userGroup.GET("", controller.ListUsers)
	userGroup.GET("/:id", controller.GetUser)
	userGroup.POST("", controller.CreateUser)
	userGroup.GET("/:id/organizations", controller.ListUserOrganizations)
}
