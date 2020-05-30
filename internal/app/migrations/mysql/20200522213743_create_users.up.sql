CREATE TABLE users (
    id bigint not null primary key,
    tgid bigint null,
    firstname varchar(100) not null,
    username varchar(100) not null
);