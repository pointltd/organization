package app

import (
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
	authController "github.com/pointltd/organization/internal/infrastructure/http/controller/auth"
	organizationController "github.com/pointltd/organization/internal/infrastructure/http/controller/organization"
	userController "github.com/pointltd/organization/internal/infrastructure/http/controller/user"
)

type controllerProvider struct {
	serviceProvider        *serviceProvider
	authController         controller.AuthController
	userController         controller.UserController
	organizationController controller.OrganizationController
}

func newControllerProvider(serviceProvider *serviceProvider) *controllerProvider {
	return &controllerProvider{
		serviceProvider: serviceProvider,
	}
}

func (c *controllerProvider) AuthController() controller.AuthController {
	if c.authController == nil {
		c.authController = authController.NewAuthController(
			c.serviceProvider.AuthenticateUserUseCase(),
			c.serviceProvider.CreateUserUseCase(),
		)
	}

	return c.authController
}

func (c *controllerProvider) UserController() controller.UserController {
	if c.userController == nil {
		c.userController = userController.NewUserController(
			c.serviceProvider.log,
			c.serviceProvider.CreateUserUseCase(),
			c.serviceProvider.ListUsersUseCase(),
			c.serviceProvider.ListUserOrganizationsUseCase(),
		)
	}

	return c.userController
}

func (c *controllerProvider) OrganizationController() controller.OrganizationController {
	if c.organizationController == nil {
		c.organizationController = organizationController.NewOrganizationController(
			c.serviceProvider.CreateOrganizationUseCase(),
			c.serviceProvider.log,
		)
	}

	return c.organizationController

}
