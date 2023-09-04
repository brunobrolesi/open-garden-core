package repository

import (
	"context"
	"errors"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
	"github.com/jackc/pgx/v5"
)

const (
	getSensorPeriodMeasurementsQuery = "SELECT sensor_id, value, time FROM measurements WHERE sensor_id=$1 AND time BETWEEN $2 AND $3"
	addSensorMeasurementQuery        = "INSERT INTO measurements (sensor_id, value) VALUES ($1, $2)"
)

type timeScaleSensorMeasurementRepository struct {
	conn *pgx.Conn
}

func NewTimeScaleSensorMeasurementRepository(conn *pgx.Conn) gateway.SensorMeasurementRepository {
	return timeScaleSensorMeasurementRepository{
		conn: conn,
	}
}

func (r timeScaleSensorMeasurementRepository) GetSensorPeriodMeasurements(ctx context.Context, sensorID int, userID int, from time.Time, to time.Time) (model.SensorMeasurements, error) {
	rows, err := r.conn.Query(ctx, getSensorPeriodMeasurementsQuery, sensorID, from, to)

	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		return model.SensorMeasurements{}, err
	}

	defer rows.Close()

	measurements := model.SensorMeasurements{}

	for rows.Next() {
		m := model.SensorMeasurement{}
		rows.Scan(&m.SensorID, &m.Value, &m.Time)
		measurements = append(measurements, m)
	}

	if err = rows.Err(); err != nil {
		return model.SensorMeasurements{}, err
	}

	return measurements, nil
}

func (r timeScaleSensorMeasurementRepository) AddSensorMeasurement(ctx context.Context, sensorID int, value float64) error {
	_, err := r.conn.Exec(ctx, addSensorMeasurementQuery, sensorID, value)
	return err
}
