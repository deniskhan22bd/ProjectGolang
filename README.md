# ProjectGolang

## Library

### API
- POST /books
- GET /books
- GET /books/:id
- PUT /books/:id
- DELETE /books/:id
- GET /health-check

### DB STRUCTURE
#### Table books
- id: bigserial [primary key]
- title: varchar
- author: varchar
- publishedYear: int
- created_at: timestamp
- updated_at: timestamp

#### Table users
- id: bigserial [primary key]
- name: varchar
- surname: varchar
- created_at: timestamp 
- updated_at: timestamp 

#### Table users_and_books (many-to-many)
- id: bigserial [primary key]
- user_id: int 
- book_id: int
- created_at: timestamp
- updated_at: timestamp

- Ref: Table users_and_books > users.id
- Ref: Table users_and_books > books.id
