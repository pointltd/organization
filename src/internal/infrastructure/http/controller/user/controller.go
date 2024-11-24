package user

import (
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
)

var _ def.UserController = (*controller)(nil)

type controller struct {
	createUserUseCase usecase.CreateUserUseCase
	listUsersUseCase  usecase.ListUsersUseCase
}

func NewController(createUserCase usecase.CreateUserUseCase, listUsersUseCase usecase.ListUsersUseCase) *controller {
	return &controller{
		createUserUseCase: createUserCase,
		listUsersUseCase:  listUsersUseCase,
	}
}
