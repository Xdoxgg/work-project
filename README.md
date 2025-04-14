

Table users {
  id serial [primary key]
  email varchar(255)
  password VARCHAR(60)
}

Table movies{
  id serial [primary key]
  title varchar(255)
  description varchar(255)
  year date
}

Table genres{
  id serial [primary key]
  title varchar(255)
}

Table tags{
  id serial [primary key]
  title varchar(255)
}

Table movies_to_tags{
  id serial [primary key]
  tag_id integer 
  movie_id integer
}

Table movies_to_genres{
  id serial [primary key]
  movie_id integer
  genre_id integer
}





ref: movies.id < movies_to_tags.movie_id
ref: tags.id < movies_to_tags.tag_id

ref: movies.id > movies_to_genres.movie_id
ref: genres.id > movies_to_genres.genre_id

