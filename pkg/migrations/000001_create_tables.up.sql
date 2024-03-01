create table if not exists books(
    id bigserial primary key,
    title varchar(255),
    author varchar(255),
    publishedYear int,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

create table if not exists users(
    id bigserial primary key,
    name varchar(255),
    surname varchar(255),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

create table if not exists users_and_books(
    id bigserial primary key,
    user_id int references users(id),
    book_id int references books(id),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);