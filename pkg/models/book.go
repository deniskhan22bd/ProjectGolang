package models

import (
	"context"
	"database/sql"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/jsonlog"
	"time"
)

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"publishedYear"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type BookModel struct {
	DB     *sql.DB
	Logger *jsonlog.Logger
}

func (m BookModel) GetAll() ([]Book, error) {
	query := `SELECT * FROM books`

	var books []Book
	var book Book
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		m.Logger.PrintError(err, nil)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedYear, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			m.Logger.PrintError(err, nil)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (m BookModel) Get(id int) (*Book, error) {
	query := `
		SELECT id, title, author, publishedYear, created_at, updated_at
		FROM books
		WHERE id = $1
		`

	var book Book

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedYear, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (m BookModel) Insert(book *Book) error {
	query := `
		INSERT INTO books (title, author, publishedYear) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
	`

	args := []interface{}{book.Title, book.Author, book.PublishedYear}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt)
}

func (m BookModel) Delete(id int) error {
	query := `
		DELETE FROM books
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func (m BookModel) Update(book *Book) error {
	query := `
		UPDATE books
		SET title = $1, author = $2, publishedyear = $3
		WHERE id = $4
		RETURNING updated_at
		`
	args := []interface{}{book.Title, book.Author, book.PublishedYear, book.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&book.UpdatedAt)
}
