CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  company_name VARCHAR(150) NOT NULL,
  email VARCHAR(150) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp,
  last_login TIMESTAMP
);

CREATE TABLE farms (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  owner INT NOT NULL,
  address VARCHAR(255) NOT NULL,
  active boolean NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp,
  FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE sensors (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  unit VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp
);

CREATE TABLE farm_sensor (
  id SERIAL PRIMARY KEY,
  farm_id INT NOT NULL,
  name VARCHAR(50) NOT NULL,
  sensor_model INT NOT NULL,
  description VARCHAR(150) NOT NULL,
  active boolean NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp,
  FOREIGN KEY (sensor_model) REFERENCES sensors (id),
  FOREIGN KEY (farm_id) REFERENCES farms (id)
);


insert into users(company_name, email, password, active) values('any_company','mail@mail.com','valid_pwd',true);
insert into farms(name, address, owner, active) values('any_farm', 'any_address', 1, true);
insert into sensors(name, type, unit) values('any_name', 'any_type', 'any_unit');
insert into farm_sensor(farm_id, name, sensor_model, description, active) values(1, 'any_name', 1, 'any_description', true);
