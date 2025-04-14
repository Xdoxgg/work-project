DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS movies;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(60) NOT NULL
);


CREATE TABLE IF NOT EXISTS movies(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    year DATE
);


CREATE TABLE IF NOT EXISTS genres(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);


Table movies_to_tags{
  id serial [primary key]
  tag_id INTEGER 
  movie_id INTEGER
}
