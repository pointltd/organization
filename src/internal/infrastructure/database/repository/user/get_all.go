package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = make([]entity.User, 0)

	all, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])

	if err != nil {
		r.log.Error(fmt.Sprintf("failed to map row: %s", err.Error()))
		return nil, err
	}

	for _, row := range all {
		user := r.userMapper.MapModelToEntity(row)

		users = append(users, user)
	}

	return users, nil
}
