package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "")
}
