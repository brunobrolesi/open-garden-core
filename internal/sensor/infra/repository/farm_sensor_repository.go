package repository

import (
	"context"
	"errors"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/jackc/pgx/v5"
)

const createFarmSensorQuery = "INSERT INTO farm_sensor(name, farm_id, sensor_model, description) values($1, $2, $3, $4) returning (id)"
const getFarmSensorByIdQuery = "SELECT id, name, farm_id, sensor_model, description, active FROM farm_sensor WHERE id = $1"
const getFarmSensorsByFarmIdQuery = "SELECT id, name, farm_id, sensor_model, description, active FROM farm_sensor WHERE farm_id = $1"

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

func (r farmSensorRepository) GetFarmSensorById(ctx context.Context, id int) (model.FarmSensor, error) {
	var farmSensor model.FarmSensor
	err := r.conn.QueryRow(ctx, getFarmSensorByIdQuery, id).Scan(&farmSensor.Id, &farmSensor.Name, &farmSensor.FarmId, &farmSensor.SensorModel, &farmSensor.Description, &farmSensor.Active)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.FarmSensor{}, err
	}

	return farmSensor, nil
}

func (r farmSensorRepository) GetFarmSensorsByFarmId(ctx context.Context, farmId int) (model.FarmSensors, error) {
	rows, err := r.conn.Query(ctx, getFarmSensorsByFarmIdQuery, farmId)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.FarmSensors{}, err
	}

	defer rows.Close()

	farmSensors := model.FarmSensors{}

	for rows.Next() {
		f := model.FarmSensor{}
		rows.Scan(&f.Id, &f.Name, &f.FarmId, &f.SensorModel, &f.Description, &f.Active)
		farmSensors = append(farmSensors, f)
	}

	if err = rows.Err(); err != nil {
		return model.FarmSensors{}, err
	}

	return farmSensors, nil
}
