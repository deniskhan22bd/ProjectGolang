# ProjectGolang

# Library

# API
POST /book
GET /book
GET /book/:id
PUT /book/:id
DELETE /book/:id

# DB STRUCTURE
Table books {
  id            : bigserial [primary key]
  title         : varchar
  author        : varchar
  publishedYear : int
  created_at    : timestamp
  updated_at    : timestamp
}

Table users {
  id           : bigserial [primary key]
  name         : varchar
  surname      : varchar
  created_at   : timestamp 
  updated_at   : timestamp 
}

// many-to-many
Table users_and_books {
  id           : bigserial [primary key]
  user_id      : int 
  book_id      : int
  created_at   : timestamp with time zone [NOT NULL] [DEFAULT NOW()]
  updated_at   : timestamp with time zone [NOT NULL] [DEFAULT NOW()]
}

Ref: Table users_and_books > users.id
ref: Table users_and_books > books.id

