CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TYPE sex AS ENUM ('female', 'male');

CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(50) not null unique,
    password_hash varchar(255) not null,
    user_role     user_role    not null
);

CREATE TABLE actors
(
    id       serial       not null unique,
    name     varchar(50) not null,
    sex      sex          not null,
    birthday date         not null
);

CREATE TABLE films
(
    id          serial        not null unique,
    name        varchar(50)   not null,
    description varchar(1000) not null,
    date        date          not null,
    rating      float         not null
);

CREATE TABLE films_actors
(
    id       serial                                                          not null unique,
    film_id  int references films (id) on delete cascade on update restrict  not null,
    actor_id int references actors (id) on delete cascade on update restrict not null,
    unique (film_id, actor_id)
);