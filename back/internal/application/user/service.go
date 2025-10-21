package user

import (
	"context"

	"github.com/RoundRonin/name-subject-to-change/back/internal/domain/user"
)

type Repository interface {
	Create(ctx context.Context, u *user.User) error
	GetByID(ctx context.Context, id string) (*user.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, name, email string) (*user.User, error) {
	u := user.NewUser(name, email)
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetUser(ctx context.Context, id string) (*user.User, error) {
	return s.repo.GetByID(ctx, id)
}

