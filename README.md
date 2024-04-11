# ProjectGolang

## Library


### GIVE PERMISSIONS FOR CREATED USER
-INSERT INTO users_permissions (user_id, permission_id)
-VALUES
-    ((SELECT id FROM users WHERE email = 'example@gmail.com'), (SELECT id FROM permissions WHERE code = 'books:write')),
-    ((SELECT id FROM users WHERE email = 'example@gmail.com'), (SELECT id FROM permissions WHERE code = 'books:delete'));
    
### API STRUCTURE
- GET /health-check

#### BOOKS
- POST /books
- GET /books
- GET /books/:id
- PUT /books/:id
- DELETE /books/:id

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
