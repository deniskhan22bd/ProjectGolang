create table if not exists comments(
    id bigserial primary key,
    comment text not NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
)

create table if not exists comments_books(
    comment_id bigint not NULL REFERENCES comments on DELETE CASCADE,
    book_id bigint not NULL REFERENCES books on DELETE CASCADE,
    primary key(comment_id, book_id)
)