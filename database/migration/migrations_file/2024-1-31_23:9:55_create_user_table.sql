CREATE TABLE users (
    id bigint auto_increment primary key,
    name varchar(100),
    email varchar(100),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);