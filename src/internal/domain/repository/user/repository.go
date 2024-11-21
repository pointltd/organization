package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/entity"
	def "github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"log"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db         *pgxpool.Pool
	userMapper mapper.UserMapper
}

func NewRepository(db *pgxpool.Pool, userMapper mapper.UserMapper) *repository {
	return &repository{
		db:         db,
		userMapper: userMapper,
	}
}

func (r *repository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		user, err := r.userMapper.MapRowToUser(rows)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
		log.Println("User: ", user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) Save(user entity.User) error {
	return nil
}
