package user

import (
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ def.UserController = (*controller)(nil)

type controller struct {
	log                          *slog.Logger
	createUserUseCase            usecase.CreateUserUseCase
	listUsersUseCase             usecase.ListUsersUseCase
	listUserOrganizationsUseCase usecase.ListUserOrganizationsUseCase
}

func NewUserController(
	log *slog.Logger,
	createUserCase usecase.CreateUserUseCase,
	listUsersUseCase usecase.ListUsersUseCase,
	listUserOrganizationsUseCase usecase.ListUserOrganizationsUseCase,
) *controller {
	return &controller{
		log:                          log,
		createUserUseCase:            createUserCase,
		listUsersUseCase:             listUsersUseCase,
		listUserOrganizationsUseCase: listUserOrganizationsUseCase,
	}
}
