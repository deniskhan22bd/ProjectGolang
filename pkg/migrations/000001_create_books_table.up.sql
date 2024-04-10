create table if not exists books(
    id bigserial primary key,
    title varchar(255),
    author varchar(255),
    publishedYear int,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO books (title, author, publishedyear)
VALUES
    ('The Great Gatsby', 'F. Scott Fitzgerald', 1925),
    ('To Kill a Mockingbird', 'Harper Lee', 1960),
    ('Pride and Prejudice', 'Jane Austen', 1813),
    ('1984', 'George Orwell', 1949),
    ('The Catcher in the Rye', 'J.D. Salinger', 1951),
    ('The Lord of the Rings', 'J.R.R. Tolkien', 1954),
    ('Harry Potter and the Philosopher''s Stone', 'J.K. Rowling', 1997),
    ('Moby Dick', 'Herman Melville', 1851),
    ('The Hobbit', 'J.R.R. Tolkien', 1937),
    ('Jane Eyre', 'Charlotte Brontë', 1847),
    ('The Grapes of Wrath', 'John Steinbeck', 1939),
    ('The Picture of Dorian Gray', 'Oscar Wilde', 1890),
    ('Wuthering Heights', 'Emily Brontë', 1847),
    ('Brave New World', 'Aldous Huxley', 1932),
    ('Crime and Punishment', 'Fyodor Dostoevsky', 1866),
    ('Frankenstein', 'Mary Shelley', 1818),
    ('Lord of the Flies', 'William Golding', 1954),
    ('The Scarlet Letter', 'Nathaniel Hawthorne', 1850),
    ('Anna Karenina', 'Leo Tolstoy', 1877),
    ('The Odyssey', 'Homer', -800),
    ('The Sun Also Rises', 'Ernest Hemingway', 1926),
    ('A Farewell to Arms', 'Ernest Hemingway', 1929);