package repository

import (
	"context"
	"database/sql"

	"github.com/RoundRonin/name-subject-to-change/back/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	query := `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.Name, u.Email, u.CreatedAt)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

