package usecase

import "github.com/pointltd/organization/internal/usecase/user"

type CreateUserUseCase interface {
	Execute(request user.CreateUserDto) error
}
