package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/validator"
)

type Comment struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	BookID    int    `json:"bookId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CommentModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m CommentModel) Get(id int) (*Comment, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, content, book_id, created_at, updated_at
		FROM comments
		WHERE id = $1
	`

	var comment Comment

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&comment.ID,
		&comment.Content,
		&comment.BookID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		return nil, ErrRecordNotFound
	}

	return &comment, nil
}

func (m CommentModel) GetByBookID(content string, book_id int, filters Filters) ([]*Comment, error) {
	query := fmt.Sprintf(`
		SELECT id, content, book_id, created_at, updated_at
		FROM comments
		WHERE (to_tsvector('simple', content) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND book_id = $2
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{content, book_id, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []*Comment{}

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.BookID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (m CommentModel) Insert(comment *Comment) error {
	query := `
		INSERT INTO comments (content, book_id) 
		VALUES ($1, $2) 
		RETURNING id, created_at, updated_at
	`

	args := []interface{}{comment.Content, comment.BookID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
}

func (m CommentModel) Delete(id int) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM comments
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func (m CommentModel) Update(comment *Comment) error {
	query := `
		UPDATE comments
		SET content = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING updated_at
	`
	args := []interface{}{comment.Content, comment.ID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&comment.UpdatedAt)
}

func ValidateComment(v *validator.Validator, comment *Comment) {
	v.Check(comment.Content != "", "content", "must be provided")
}
