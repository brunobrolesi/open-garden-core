package repository

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/jackc/pgx/v5"
)

const createFarmSensorQuery = "INSERT INTO farm_sensor(name, farm_id, sensor_model, description) values($1, $2, $3) returning (id)"

type farmSensorRepository struct {
	conn *pgx.Conn
}

func NewPostgresFarmSensorRepository(conn *pgx.Conn) gateway.FarmSensorRepository {
	return &farmSensorRepository{
		conn: conn,
	}
}

func (r farmSensorRepository) CreateFarmSensor(ctx context.Context, farmSensor model.FarmSensor) (model.FarmSensor, error) {
	err := r.conn.QueryRow(ctx, createFarmSensorQuery, farmSensor.Name, farmSensor.FarmId, farmSensor.SensorModel, farmSensor.Description).Scan(&farmSensor.Id)
	return farmSensor, err
}
