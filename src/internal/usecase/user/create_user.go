package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pointltd/organization/internal/data"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
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
	id, err := uuid.NewV7()
	if err != nil {
		log.Fatal(fmt.Sprintf("error generating user UUID: %v\n", err))
	}

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

	user := entity.User{
		ID:        id.String(),
		Info:      info,
		Contacts:  contacts,
		Timestamp: timestamps,
	}

	return user, nil
}
