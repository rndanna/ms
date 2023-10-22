CREATE SCHEMA engines;

CREATE TABLE engines.engines
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
    description TEXT
);


INSERT INTO cars.engines (name,  description) VALUES ('engine1', 'description1');
INSERT INTO cars.engines (name,  description) VALUES ('engine2', 'description2');
INSERT INTO cars.engines (name,  description) VALUES ('engine3', 'description3');

