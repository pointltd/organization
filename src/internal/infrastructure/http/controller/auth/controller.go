package auth

import (
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
)

var _ def.AuthController = (*controller)(nil)

type controller struct {
	authenticateUserUseCase usecase.AuthenticateUserUseCase
}

func NewAuthController(authenticateUserUseCase usecase.AuthenticateUserUseCase) *controller {
	return &controller{
		authenticateUserUseCase: authenticateUserUseCase,
	}
}
