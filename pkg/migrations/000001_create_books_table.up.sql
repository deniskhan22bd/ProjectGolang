create table if not exists books(
    id bigserial primary key,
    title varchar(255),
    author varchar(255),
    publishedYear int,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);