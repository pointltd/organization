package route

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/infrastructure/controller"
)

func RegisterUserRoutes(group *echo.Group, controller controller.UserController) {
	userGroup := group.Group("/users")
	userGroup.GET("", controller.ListUsers)
	userGroup.GET("/:id", controller.GetUser)
	//artistsGroup.GET("", controllers.ListArtists)
	//artistsGroup.GET("/:id", controllers.GetArtist)
	//artistsGroup.POST("", controllers.CreateArtist)
	//artistsGroup.PUT("/:id", controllers.UpdateArtist)
	//artistsGroup.DELETE("/:id", controllers.DeleteArtist)
}
