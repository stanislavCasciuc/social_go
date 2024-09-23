package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	userName  string `json:"username"`
	email     string `json:"email"`
	password  string `json:"-"`
	createdAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`
	err := s.db.QueryRowContext(ctx, query,
		u.userName,
		u.email,
		u.password,
	).Scan(&u.ID, &u.createdAt)
	if err != nil {
		return err
	}
	return nil
}
