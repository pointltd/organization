package auth

import (
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ def.AuthController = (*controller)(nil)

type controller struct {
	log                     *slog.Logger
	createUserUseCase       usecase.CreateUserUseCase
	authenticateUserUseCase usecase.AuthenticateUserUseCase
}

func NewAuthController(
	authenticateUserUseCase usecase.AuthenticateUserUseCase,
	createUserUseCase usecase.CreateUserUseCase,
) *controller {
	return &controller{
		authenticateUserUseCase: authenticateUserUseCase,
		createUserUseCase:       createUserUseCase,
	}
}
