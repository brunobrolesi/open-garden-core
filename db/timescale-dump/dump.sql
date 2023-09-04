CREATE TABLE measurements (
  time         TIMESTAMPTZ       NOT NULL default current_timestamp,
  sensor_id    DOUBLE PRECISION  NULL,
  value        DOUBLE PRECISION  NULL
);

CREATE EXTENSION IF NOT EXISTS timescaledb;

SELECT create_hypertable('measurements', 'time');

DO $$ 
DECLARE 
    startTime TIMESTAMP := CURRENT_TIMESTAMP - INTERVAL '1 day';
    intervalMinutes INT := 5;
    maxValue INT := 40;
    i INT;
BEGIN
    FOR i IN 0..49 LOOP
        EXECUTE 'INSERT INTO measurements(sensor_id, value, time) VALUES(1, ' || 
            (RANDOM() * (maxValue + 1))::INT || ', ''' || 
            (startTime + (intervalMinutes * i || ' minutes')::INTERVAL) || ''');';
    END LOOP;
END $$;


