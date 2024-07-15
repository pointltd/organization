package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/internal/controllers"
)

func RegisterArtistRoutes(group *echo.Group) {
	artistsGroup := group.Group("/artists")
	artistsGroup.GET("", controllers.ListArtists)
	artistsGroup.GET("/:id", controllers.GetArtist)
	artistsGroup.POST("", controllers.CreateArtist)
	artistsGroup.PUT("/:id", controllers.UpdateArtist)
	artistsGroup.DELETE("/:id", controllers.DeleteArtist)
}
