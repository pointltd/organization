package auth

import (
	"github.com/pointltd/organization/internal/config"
	def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ def.AuthController = (*controller)(nil)

type controller struct {
	log                     *slog.Logger
	createUserUseCase       usecase.CreateUserUseCase
	authenticateUserUseCase usecase.AuthenticateUserUseCase
	config                  config.AppConfig
}

func NewAuthController(
	log *slog.Logger,
	authenticateUserUseCase usecase.AuthenticateUserUseCase,
	createUserUseCase usecase.CreateUserUseCase,
	config config.AppConfig,
) *controller {
	return &controller{
		log:                     log,
		authenticateUserUseCase: authenticateUserUseCase,
		createUserUseCase:       createUserUseCase,
		config:                  config,
	}
}
