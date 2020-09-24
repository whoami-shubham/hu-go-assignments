create table if not exists users(
    id serial primary key,
    reputation integer,
    creation_date VARCHAR(100),
    display_name VARCHAR(50),
    last_access_date VARCHAR(100),
    website_url VARCHAR(500),
    location VARCHAR(100),
    about_me VARCHAR(100000),
    views integer,
    upvotes integer,
    downvotes integer,
    account_id integer,
    profile_image_url varchar(500)
);

create table if not exists posts(
    id serial primary key,
    post_type_id integer,
    score integer,
    view_count integer,
    tags VARCHAR(500),
    answer_count integer,
    comment_count integer,
    favourite_count integer,
    creation_date VARCHAR(100),
    body VARCHAR(100000),
    closed_date VARCHAR(100),
    accepted_answer_id integer DEFAULT -1,
    parent_id integer DEFAULT -1,
    owner_user_id integer REFERENCES users(id),
    owner_display_name varchar(200),
    last_editor_user_id integer REFERENCES users(id),
    last_editor_display_name varchar(200),
    last_edit_date varchar(200),
    last_activity_date varchar(200),
    title varchar(500),
    community_owned_date varchar(500)
);

create table if not exists badges(
    id serial primary key,
    user_id integer REFERENCES users(id),
    date VARCHAR(100),
    name VARCHAR(50),
    class integer,
    tagBased VARCHAR(50)
);

create table if not exists comments(
    id serial primary key,
    post_id integer REFERENCES posts(id),
    creation_date VARCHAR(100),
    text VARCHAR(100000),
    user_id integer REFERENCES users(id),
    score integer,
    user_display_name varchar(100)
);

create table if not exists post_history(
    id serial primary key,
    post_history_type_id integer,
    post_id integer REFERENCES posts(id),
    revision_guid VARCHAR(100),
    creation_date VARCHAR(100),
    text VARCHAR(100000),
    user_id integer REFERENCES users(id),
    user_display_name varchar(100),
    comment varchar(500)
);

create table if not exists post_link(
    id serial primary key,
    post_id integer REFERENCES posts(id),
    creation_date VARCHAR(100),
    related_post_id integer REFERENCES posts(id),
    link_type_id integer
);

create table if not exists tags(
    id serial primary key,
    tag_name VARCHAR(50),
    count integer,
    excerpt_post_id integer REFERENCES posts(id),
    wiki_post_id integer REFERENCES posts(id)
);

create table if not exists votes(
    id serial primary key,
    post_id integer REFERENCES posts(id),
    creation_date VARCHAR(100),
    vote_type_id integer,
    user_id integer REFERENCES users(id),
    bounty_amount integer
);

create table if not exists customer(
    id serial primary key,
    username VARCHAR(100) UNIQUE,
    password VARCHAR(10),
    email VARCHAR(100)
);