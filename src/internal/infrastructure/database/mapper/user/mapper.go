package user

import (
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	_def "github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

var _ _def.UserMapper = (*userMapper)(nil)

type userMapper struct{}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) MapModelToEntity(model model.User) entity.User {
	var timestamp = entity.Timestamp{}

	if model.CreatedAt.Valid {
		timestamp.CreatedAt = &model.CreatedAt.Time
	}

	if model.UpdatedAt.Valid {
		timestamp.UpdatedAt = &model.UpdatedAt.Time
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

	return user
}

func (m *userMapper) MapEntityToArg(user entity.User) pgx.NamedArgs {
	args := pgx.NamedArgs{
		"id":         user.ID,
		"password":   user.Password,
		"first_name": user.Info.FirstName,
		"email":      user.Contacts.Email,
	}

	if user.Info.LastName != nil {
		args["last_name"] = *user.Info.LastName
	}

	if user.Contacts.Phone != nil {
		args["phone"] = *user.Contacts.Phone
	}

	if user.Timestamp.CreatedAt != nil {
		args["created_at"] = *user.Timestamp.CreatedAt
	}

	if user.Timestamp.UpdatedAt != nil {
		args["updated_at"] = *user.Timestamp.UpdatedAt
	}

	if user.Timestamp.DeletedAt != nil {
		args["deleted_at"] = *user.Timestamp.DeletedAt
	}

	if user.UserStamp.CreatedById != nil {
		args["created_by_id"] = *user.UserStamp.CreatedById
	}

	if user.UserStamp.UpdatedById != nil {
		args["updated_by_id"] = *user.UserStamp.UpdatedById
	}

	if user.UserStamp.DeletedById != nil {
		args["deleted_by_id"] = *user.UserStamp.DeletedById
	}

	return args
}
