CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  company_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  last_login TIMESTAMP
);

CREATE TABLE farms (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  owner INT NOT NULL,
  address VARCHAR(255) NOT NULL,
  active boolean NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE sensors (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  unit VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE farm_sensor (
  id SERIAL PRIMARY KEY,
  farm_id INT NOT NULL,
  sensor_model INT NOT NULL,
  description VARCHAR(255) NOT NULL,
  active boolean NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (sensor_model) REFERENCES sensors (id),
  FOREIGN KEY (farm_id) REFERENCES farms (id)
);

