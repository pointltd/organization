package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	ListUsers(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
}
