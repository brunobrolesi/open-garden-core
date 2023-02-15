package repository

import (
	"context"
	"errors"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/jackc/pgx/v5"
)

const getSensorByIdQuery = "SELECT id, name, type, unit FROM sensors WHERE id=$1"

type sensorRepository struct {
	conn *pgx.Conn
}

func NewPostgresSensorRepository(conn *pgx.Conn) gateway.SensorRepository {
	return &sensorRepository{
		conn: conn,
	}
}

func (r sensorRepository) GetSensorById(ctx context.Context, id int) (model.Sensor, error) {
	sensor := model.Sensor{}
	err := r.conn.QueryRow(ctx, getSensorByIdQuery, id).Scan(&sensor.Id, &sensor.Name, &sensor.Type, &sensor.Unit)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.Sensor{}, err
	}

	return sensor, nil
}
