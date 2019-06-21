DROP TABLE IF EXISTS users;
CREATE TABLE users(id serial PRIMARY KEY, email VARCHAR(250), name VARCHAR(250), age integer);