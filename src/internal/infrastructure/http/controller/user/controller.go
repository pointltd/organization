package user

import (
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ def.UserController = (*controller)(nil)

type controller struct {
	log               *slog.Logger
	createUserUseCase usecase.CreateUserUseCase
	listUsersUseCase  usecase.ListUsersUseCase
}

func NewController(createUserCase usecase.CreateUserUseCase, listUsersUseCase usecase.ListUsersUseCase) *controller {
	return &controller{
		createUserUseCase: createUserCase,
		listUsersUseCase:  listUsersUseCase,
	}
}
