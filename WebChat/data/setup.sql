drop table posts;
drop table threads;
drop table sessions;
drop table users;


create table users (
    id uuid default gen_random_uuid() primary key,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp
);

create table sessions(
    id uuid default gen_random_uuid() primary key,
    user_id uuid references users(id) not null unique,
    email varchar(255),
    created_at timestamp default current_timestamp
);


create table threads (
    id uuid default gen_random_uuid() primary key,
    user_id uuid references users(id) not null,
    title varchar(255) not null,
    created_at timestamp default current_timestamp
);

create table posts (
    id uuid default gen_random_uuid() primary key,
    thread_id uuid references threads(id) not null,
    user_id uuid references users(id) not null,
    content text not null,
    created_at timestamp default current_timestamp
);