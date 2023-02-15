package repository

import (
	"context"
	"errors"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/jackc/pgx/v5"
)

const getFarmByIdAndUserIdQuery = "SELECT id, name, address, owner, active FROM farms WHERE id=$1 AND owner=$2"

type farmRepository struct {
	conn *pgx.Conn
}

func NewPostgresFarmRepository(conn *pgx.Conn) gateway.FarmRepository {
	return &farmRepository{
		conn: conn,
	}
}

func (r farmRepository) GetFarmByIdAndUserId(ctx context.Context, id int, userId int) (model.Farm, error) {
	farm := model.Farm{}
	err := r.conn.QueryRow(ctx, getFarmByIdAndUserIdQuery, id, userId).Scan(&farm.Id, &farm.Name, &farm.Address, &farm.Owner, &farm.Active)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.Farm{}, err
	}

	return farm, nil
}
