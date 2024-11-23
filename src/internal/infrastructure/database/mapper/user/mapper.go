package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	_def "github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

var _ _def.UserMapper = (*userMapper)(nil)

type userMapper struct{}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) MapModelToEntity(model model.User) (entity.User, error) {
	var timestamp = entity.Timestamp{
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	if model.DeletedAt.Valid {
		timestamp.DeletedAt = &model.DeletedAt.Time
	}

	var contactInfo = entity.ContactInfo{
		Email: model.Email,
	}

	if model.Phone.Valid {
		contactInfo.Phone = &model.Phone.String
	}

	var userStamp = entity.UserStamp{}

	if model.CreatedById.Valid {
		userStamp.CreatedById = &model.CreatedById.String
	}
	if model.UpdatedById.Valid {
		userStamp.UpdatedById = &model.UpdatedById.String
	}
	if model.DeletedById.Valid {
		userStamp.DeletedById = &model.DeletedById.String
	}

	var user = entity.User{
		ID:       model.ID,
		Password: model.Password,
		Info: entity.UserInfo{
			FirstName: model.FirstName,
		},
		Contacts:  contactInfo,
		Timestamp: timestamp,
		UserStamp: userStamp,
	}

	if model.LastName.Valid {
		user.Info.LastName = &model.LastName.String
	}

	return user, nil
}
