CREATE TABLE users (
    id bigint AUTO_INCREMENT UNIQUE not null primary key,
    tgid bigint null,
    firstname varchar(100) not null,
    username varchar(100) not null
);