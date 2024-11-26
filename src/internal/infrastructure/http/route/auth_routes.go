package route

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
)

func RegisterAuthRoutes(group *echo.Group, controller controller.AuthController) {
	authGroup := group.Group("")
	authGroup.POST("/login", controller.Login)
}
