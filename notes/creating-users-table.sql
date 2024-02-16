drop table if exists users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT UNIQUE
);
