-- THIS IS SCRIPT FOR CREATING USERS TABLE
CREATE TABLE IF NOT EXISTS users (
   uuid UUID PRIMARY KEY NOT NULL,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL,
   password TEXT NOT NULL,
   role VARCHAR(10) NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_by VARCHAR(255),
   deleted_by VARCHAR(255),
   CONSTRAINT unique_email UNIQUE (email)
);

-- THIS IS SCRIPT FOR CREATING TOKENS TABLE
CREATE TABLE IF NOT EXISTS tokens (
   uuid UUID PRIMARY KEY NOT NULL,
   user_uuid UUID NOT NULL,
   token TEXT NOT NULL,
   is_active BOOLEAN NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_by VARCHAR(255),
   FOREIGN KEY (user_uuid) REFERENCES users(uuid)
);

-- THIS IS SCRIPT FOR CREATING MOVIES TABLE
CREATE TABLE IF NOT EXISTS movies (
   uuid UUID PRIMARY KEY NOT NULL,
   title VARCHAR(255) NOT NULL,
   description TEXT NOT NULL,
   duration INTEGER NOT NULL,
   artists UUID[] NOT NULL,
   genres UUID[] NOT NULL,
   url VARCHAR(255) NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_by VARCHAR(255),
   deleted_by VARCHAR(255),
   CONSTRAINT unique_title_movie UNIQUE (title)
);

-- THIS IS SCRIPT FOR CREATING GENRES TABLE
CREATE TABLE IF NOT EXISTS genres (
    uuid UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT unique_name_genre UNIQUE (name)
);

-- THIS IS SCRIPT FOR CREATING ARTISTS TABLE
CREATE TABLE IF NOT EXISTS artists (
    uuid UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT unique_name_artist UNIQUE (name)
);

-- THIS IS SCRIPT FOR CREATING VOTERS TABLE
CREATE TABLE IF NOT EXISTS voters (
    uuid UUID PRIMARY KEY NOT NULL,
    movie_uuid UUID NOT NULL,
    user_uuid UUID NOT NULL,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by VARCHAR(255),
    deleted_by VARCHAR(255),
    FOREIGN KEY (movie_uuid) REFERENCES movies(uuid),
    FOREIGN KEY (user_uuid) REFERENCES users(uuid),
    UNIQUE (movie_uuid, user_uuid)
);

INSERT INTO genres (uuid, name) VALUES
(gen_random_uuid(), 'Action'),
(gen_random_uuid(), 'Comedy'),
(gen_random_uuid(), 'Crime'),
(gen_random_uuid(), 'Drama'),
(gen_random_uuid(), 'Horror'),
(gen_random_uuid(), 'Romance'),
(gen_random_uuid(), 'Sci-Fi'),
(gen_random_uuid(), 'Thriller'),
(gen_random_uuid(), 'Fantasy'),
(gen_random_uuid(), 'Animation'),
(gen_random_uuid(), 'Documentary');

INSERT INTO artists (uuid, name) VALUES
 (gen_random_uuid(), 'Leonardo DiCaprio'),
 (gen_random_uuid(), 'Scarlett Johansson'),
 (gen_random_uuid(), 'Tom Hanks'),
 (gen_random_uuid(), 'Natalie Portman'),
 (gen_random_uuid(), 'Denzel Washington'),
 (gen_random_uuid(), 'Meryl Streep'),
 (gen_random_uuid(), 'Robert Downey Jr.'),
 (gen_random_uuid(), 'Angelina Jolie'),
 (gen_random_uuid(), 'Brad Pitt'),
 (gen_random_uuid(), 'Chris Hemsworth');