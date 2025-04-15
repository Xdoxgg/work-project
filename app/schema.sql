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




CREATE TABLE IF NOT EXISTS movies_to_tags(
  id SERIAL PRIMARY KEY,
  tag_id INTEGER,
  movie_id INTEGER,
  FOREIGN KEY tag_id REFERENCESE tags.id,
  FOREIGN KEY movie_id REFERENCESE movies.id
);

CREATE TABLE IF NOT EXISTS movies_to_genres(
  id SERIAL PRIMARY KEY,
  movie_id INTEGER,
  genre_id INTEGER,
);
