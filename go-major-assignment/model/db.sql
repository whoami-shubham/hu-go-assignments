create table if not exists users(
    id serial primary key,
    name VARCHAR(50),
    password VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    role VARCHAR(50),
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp
);

create table if not exists projects(
    id serial primary key,
    name VARCHAR(20),
    created_by integer,
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp
);

create table if not exists issues(
    id serial primary key,
    title VARCHAR(50),
    description Text,
    type VARCHAR(5),
    assignee integer,
    reporter integer,
    status VARCHAR(20),
    project integer,
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp
);

create table if not exists issues_log(
    id serial primary key,
    updated_feild VARCHAR(50),
    previous_value Text,
    new_value Text,
    issue_id integer,
    updated_on timestamp default current_timestamp
);

create table if not exists comments(
    id serial primary key,
    author integer,
    text Text,
    issue_id integer,
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp
);

create table if not exists watchers(
    id serial primary key,
    user_id integer,
    issue_id integer
);