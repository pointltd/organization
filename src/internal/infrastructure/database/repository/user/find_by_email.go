package user

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) FindByEmail(email string) (*entity.User, error) {
	row, err := r.db.Query(context.Background(), "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	userModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	user := r.userMapper.MapModelToEntity(userModel)

	return &user, nil
}
