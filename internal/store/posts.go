package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64
	content   string
	title     string
	userID    int64
	tags      []string
	createdAt string
	updatedAt string
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.content,
		post.title,
		post.userID,
		pq.Array(post.tags),
	).Scan(&post.ID, &post.createdAt, &post.updatedAt)
	if err != nil {
		return err

	}

	return nil
}
