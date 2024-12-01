package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) Save(user entity.User) (entity.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		r.log.Error(fmt.Sprintf("error generating user UUID: %v", err))
		return user, err
	}

	user.Id = id.String()

	query :=
		`INSERT INTO users (id, password, first_name, last_name, email, phone, created_at, updated_at, deleted_at) 
		VALUES (@id, @password, @first_name, @last_name, @email, @phone, @created_at, @updated_at, @deleted_at) 
		RETURNING *`

	args := r.userMapper.MapEntityToArg(user)

	row, err := r.db.Query(context.Background(), query, args)

	if err != nil {
		return user, err
	}

	defer row.Close()

	userModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.User])
	if err != nil {
		return user, err
	}

	user = r.userMapper.MapModelToEntity(userModel)

	return user, nil
}
