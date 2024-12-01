package controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type UserController interface {
	ListUsers(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	ListUserOrganizations(ctx echo.Context) error
}

type OrganizationController interface {
	CreateOrganization(ctx echo.Context) error
	ListOrganizationPoints(ctx echo.Context) error
}

type PointController interface {
	CreatePoint(ctx echo.Context) error
}
