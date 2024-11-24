package user

import (
	"fmt"
	"github.com/pointltd/organization/internal/data"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
	"github.com/pointltd/organization/pkg/password"
	"log"
	"time"
)

var _ def.CreateUserUseCase = (*createUserUseCase)(nil)

type createUserUseCase struct {
	userRepository repository.UserRepository
}

func NewUseCase(userRepository repository.UserRepository) *createUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}

func (u createUserUseCase) Execute(dto data.CreateUserDTO) (entity.User, error) {
	info := entity.UserInfo{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	contacts := entity.ContactInfo{
		Email: dto.Email,
	}

	timestamps := entity.Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPassword, err := password.HashPassword(dto.Password)

	if err != nil {
		log.Fatal(fmt.Sprintf("error hashing password: %v\n", err))
	}

	user := entity.User{
		Password:  hashedPassword,
		Info:      info,
		Contacts:  contacts,
		Timestamp: timestamps,
	}

	err = u.userRepository.Save(user)

	return user, nil
}
