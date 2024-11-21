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
	var firstName, phone, createdById, updatedById, deletedById sql.NullString
	var deletedAt sql.NullTime

	err := rows.Scan(
		&user.ID,
		&firstName,
		&user.Info.LastName,
		&user.Contacts.Email,
		&user.Password,
		&phone,
		&user.Timestamp.CreatedAt,
		&user.Timestamp.UpdatedAt,
		&deletedAt,
		&createdById,
		&updatedById,
		&deletedById,
	)
	if err != nil {
		return user, err
	}

	if phone.Valid {
		user.Contacts.Phone = phone.String
	}
	if firstName.Valid {
		user.Info.FirstName = firstName.String
	}
	if deletedAt.Valid {
		user.Timestamp.DeletedAt = &deletedAt.Time
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
