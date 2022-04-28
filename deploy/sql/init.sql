CREATE schema xm_db;
SET search_path to xm_db, public;
CREATE TABLE IF NOT EXISTS users
(
    id            uuid default gen_random_uuid() PRIMARY KEY,
    username      varchar(100) not null unique,
    password_hash varchar(100) not null
);

CREATE TABLE IF NOT EXISTS countries
(
    id         serial PRIMARY KEY,
    name       varchar not null
);

CREATE TABLE IF NOT EXISTS companies
(
    id         serial PRIMARY KEY,
    name       varchar not null,
    code       integer not null,
    country_id     serial    not null references countries (id),
    website    varchar not null,
    phone      varchar not null,
    created_at integer default null,
    updated_at integer default null
);