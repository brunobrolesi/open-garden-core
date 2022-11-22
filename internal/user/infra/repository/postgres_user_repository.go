package repository

import (
	"context"
	"errors"

	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/jackc/pgx/v5"
)

const (
	createUserQuery     = "insert into users(company_name, email, password, active) values($1, $2, $3, $4) returning (id)"
	getUserByEmailQuery = "select id, company_name, email, password, active from users where email=$1"
	checkEmailInUse     = "select email from users where email=$1"
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

	if shared.IsPostgreSqlError(err, shared.POSTGRESQL_UNIQUE_VIOLATION_CODE) {
		return model.User{}, model.ErrEmailInUse
	}

	return user, err
}

func (r postgresUserRepository) GetUserByEmail(email string, ctx context.Context) (model.User, error) {
	user := model.User{}
	err := r.conn.QueryRow(ctx, getUserByEmailQuery, email).Scan(&user.Id, &user.CompanyName, &user.Email, &user.Password, &user.Active)

	if !errors.Is(err, pgx.ErrNoRows) {
		return user, err
	}

	return user, nil
}

func (r postgresUserRepository) IsEmailInUse(email string, ctx context.Context) (bool, error) {
	var e string
	err := r.conn.QueryRow(ctx, checkEmailInUse, email).Scan(&e)

	if err != nil {
		return false, err
	}

	if e != "" {
		return true, nil
	}

	return false, nil
}
