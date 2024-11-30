package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	def "github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"log/slog"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db         *pgxpool.Pool
	userMapper mapper.UserMapper
	log        *slog.Logger
}

func NewUserRepository(db *pgxpool.Pool, userMapper mapper.UserMapper, log *slog.Logger) *repository {
	return &repository{
		db:         db,
		userMapper: userMapper,
		log:        log,
	}
}
