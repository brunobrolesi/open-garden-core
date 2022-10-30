package repository

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/jackc/pgx/v5"
)

const (
	createUserQuery = "insert into users(company_name, email, password, active) values($1, $2, $3, $4) returning (id)"
)

type postgresUserRepository struct {
	conn *pgx.Conn
}

func NewPostgresUserRepository(conn *pgx.Conn) gateway.UserRepository {
	return postgresUserRepository{
		conn: conn,
	}
}

func (r postgresUserRepository) CreateUser(user model.User, ctx context.Context) (model.User, error) {
	err := r.conn.QueryRow(ctx, createUserQuery, user.CompanyName, user.Email, user.Password, user.Active).Scan(&user.Id)
	return user, err
}
