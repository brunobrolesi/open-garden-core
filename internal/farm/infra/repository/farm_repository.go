package repository

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/jackc/pgx/v5"
)

const (
	createFarmQuery = "insert into farms(name, address, owner, active) values($1, $2, $3, $4) returning (id)"
)

type postgresFarmRepository struct {
	conn *pgx.Conn
}

func NewPostgresFarmRepository(conn *pgx.Conn) gateway.FarmRepository {
	return postgresFarmRepository{
		conn: conn,
	}
}

func (r postgresFarmRepository) CreateFarm(farm model.Farm, ctx context.Context) (model.Farm, error) {
	err := r.conn.QueryRow(ctx, createFarmQuery, farm.Name, farm.Address, farm.Owner, farm.Active).Scan(&farm.Id)
	return farm, err
}
