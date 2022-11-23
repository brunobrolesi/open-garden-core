package repository

import (
	"context"
	"errors"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/jackc/pgx/v5"
)

const (
	createFarmQuery      = "insert into farms(name, address, owner, active) values($1, $2, $3, $4) returning (id)"
	getAllUserFarmsQuery = "SELECT id, name, address, owner, active FROM farms WHERE owner=$1"
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

func (r postgresFarmRepository) GetFarmsByUserId(userId int, ctx context.Context) (model.Farms, error) {
	rows, err := r.conn.Query(ctx, getAllUserFarmsQuery, userId)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.Farms{}, err
	}

	defer rows.Close()

	farms := model.Farms{}

	for rows.Next() {
		f := model.Farm{}
		rows.Scan(&f.Id, &f.Name, &f.Address, &f.Owner, &f.Active)
		farms = append(farms, f)
	}

	if err = rows.Err(); err != nil {
		return model.Farms{}, err
	}

	return farms, nil
}
