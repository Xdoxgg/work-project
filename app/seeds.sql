INSERT INTO users (email, password)
VALUES ('alice@example.com', '$2a$10$T.UX3A759amfWh4AwZbuau2nwPpCLDfTiJ6ekQ7nUnWKkCO19j.pa'),   -- password1
       ('bob@example.com', '$2a$10$JCtmSEyXIEMdutqI/HBij.zUuJZWqGrvVK01OQPH5YBu4HVqx3rYC'),     -- password2
       ('charlie@example.com', '$2a$10$xfsZdLYLxZcg1Iei5SWk0.Rxm/LCE4QOJCvUyZZ4Fvs.LtB7H2FkK'), -- password3
       ('dave@example.com', '$2a$10$UKP43aIeKRKjOR.dxajsLOdiVbd57PL.22DnlEh1XtVgdXp3xiBOS'),    -- password4
       ('eve@example.com', '$2a$10$YV/FX4FXyNY539IgZ7GbqO0XKT1YciuJx47jtXhI7ZV4yWsaFdUUi'); -- password5


INSERT INTO movies (title, description, year)
VALUES ('Inception',
        'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.',
        '2010-07-16'),
       ('The Matrix',
        'A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.',
        '1999-03-31'),
       ('Interstellar',
        'A team of explorers travel through a wormhole in space in an attempt to ensure humanity''s survival.',
        '2014-11-07'),
       ('The Shawshank Redemption',
        'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
        '1994-09-23'),
       ('The Godfather',
        'An organized crime dynasty''s aging patriarch transfers control of his clandestine empire to his reluctant son.',
        '1972-03-24');

INSERT INTO genres (title)
VALUES ('Action'),
       ('Drama'),
       ('Sci-Fi'),
       ('Thriller'),
       ('Crime'),
       ('Adventure');

-- Сиды для тегов
INSERT INTO tags (title)
VALUES ('Mind-bending'),
       ('Classic'),
       ('Dystopian'),
       ('Inspirational'),
       ('Friendship'),
       ('Redemption');

-- Сиды для связи фильмов и тегов
INSERT INTO movies_to_tags (tag_id, movie_id)
VALUES (1, 1), -- Mind-bending for Inception
       (2, 4), -- Classic for The Shawshank Redemption
       (3, 2), -- Dystopian for The Matrix
       (4, 1), -- Inspirational for Inception
       (5, 4); -- Friendship for The Shawshank Redemption

-- Сиды для связи фильмов и жанров
INSERT INTO movies_to_genres (movie_id, genre_id)
VALUES (1, 3), -- Inception is Sci-Fi
       (2, 1), -- The Matrix is Action
       (3, 3), -- Interstellar is Sci-Fi
       (4, 2), -- The Shawshank Redemption is Drama
       (5, 2); -- The Godfather is Drama