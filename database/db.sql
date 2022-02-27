DROP DATABASE shortener;
CREATE DATABASE shortener;

CREATE TABLE IF NOT EXISTS shortener.users(
	id              INTEGER AUTO_INCREMENT,
    name            VARCHAR(80) NOT NULL,
    email           VARCHAR(255) NOT NULL, UNIQUE,
    password        VARCHAR(1000) NOT NULL,

    CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS shortener.sessions(
    id              INTEGER AUTO_INCREMENT,
    user_id         INTEGER NOT NULL,UNIQUE
    token           VARCHAR(36) NOT NULL,UNIQUE

    CONSTRAINT sessions_pk PRIMARY KEY(id),
    CONSTRAINT sessions_fk FOREIGN KEY(user_id) REFERENCES shortener.users(id)
);

CREATE TABLE IF NOT EXISTS shortener.urls(
    id               INTEGER AUTO_INCREMENT,
    full_url         varchar(80) NOT NULL,
    shortened_url    VARCHAR(10) NOT NULL,
    generated_by     INTEGER NOT NULL,UNIQUE,

    CONSTRAINT urls_pk PRIMARY KEY(id),
    CONSTRAINT urls_fk FOREIGN KEY(generated_by) REFERENCES shortener.users(id)
);