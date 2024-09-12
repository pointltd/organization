package user

import (
	def "github.com/pointltd/organization/internal/infrastructure/controller"
	"github.com/pointltd/organization/internal/usecase"
)

var _ def.UserController = (*controller)(nil)

type controller struct {
	createUserCase usecase.CreateUserUseCase
}

func NewController(useCase usecase.CreateUserUseCase) *controller {
	return &controller{
		createUserCase: useCase,
	}
}
