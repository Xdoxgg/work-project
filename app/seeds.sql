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
        'A team of explorers travel through a wormhole in space in an attempt to ensure humanity''s survival.', '2014-11-07'),
       ('The Shawshank Redemption',
        'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
        '1994-09-23'),
       ('The Godfather',
        'An organized crime dynasty''s aging patriarch transfers control of his clandestine empire to his reluctant son.',
        '1972-03-24');
