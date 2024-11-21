package user

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	_def "github.com/pointltd/organization/internal/infrastructure/database/mapper"
)

var _ _def.UserMapper = (*userMapper)(nil)

type userMapper struct{}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) MapRowToUser(rows pgx.Rows) (entity.User, error) {
	var user entity.User
	var firstName, createdById, updatedById, deletedById sql.NullString

	err := rows.Scan(
		&user.ID,
		&user.Password,
		&user.Contacts.Email,
		&user.Contacts.Phone,
		&firstName,
		&user.Info.LastName,
		&user.Timestamp.CreatedAt,
		&user.Timestamp.UpdatedAt,
		&user.Timestamp.DeletedAt,
		&createdById,
		&updatedById,
		&deletedById,
	)
	if err != nil {
		return user, err
	}

	if firstName.Valid {
		user.Info.FirstName = firstName.String
	}
	if createdById.Valid {
		user.UserStamp.CreatedById = createdById.String
	}
	if updatedById.Valid {
		user.UserStamp.UpdatedById = updatedById.String
	}
	if deletedById.Valid {
		user.UserStamp.DeletedById = deletedById.String
	}

	return user, nil
}
