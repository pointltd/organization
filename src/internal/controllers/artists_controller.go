package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pointltd/organization/pkg/models"
)

var artists = []models.Artist{
	{ID: 1, Name: "Leonardo da Vinci", Genre: "Renaissance", Country: "Italy"},
	{ID: 2, Name: "Pablo Picasso", Genre: "Cubism", Country: "Spain"},
}

func ListArtists(c echo.Context) error {
	return c.JSON(http.StatusOK, artists)
}

func GetArtist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, artist := range artists {
		if artist.ID == id {
			return c.JSON(http.StatusOK, artist)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Artist not found"})
}

func CreateArtist(c echo.Context) error {
	artist := models.Artist{}
	if err := c.Bind(&artist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	artist.ID = len(artists) + 1
	artists = append(artists, artist)
	return c.JSON(http.StatusCreated, artist)
}

func UpdateArtist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, artist := range artists {
		if artist.ID == id {
			updatedArtist := models.Artist{}
			if err := c.Bind(&updatedArtist); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
			}
			updatedArtist.ID = id
			artists[index] = updatedArtist
			return c.JSON(http.StatusOK, updatedArtist)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Artist not found"})
}

func DeleteArtist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, artist := range artists {
		if artist.ID == id {
			artists = append(artists[:index], artists[index+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Artist not found"})
}
