package usersService

import (
	users "project01/app/internal/db/users"
	"context"
	"fmt"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo *users.Repo
}

type CreateUserInput struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
}

func New(repo *users.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]users.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
           return nil, fmt.Errorf("user not found: %w", err)
        }
		return nil, fmt.Errorf("service: failed to get users: %w", err)
	}
	return users, nil
}

func (s *Service) GetById(ctx context.Context, id int) (*users.User, error) {
		
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
           return nil, fmt.Errorf("user not found: %w", err)
        }
		return nil, fmt.Errorf("service: failed to get user: %w", err)
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, input CreateUserInput) (*users.User, error) {
    if input.FirstName == "" || input.Email == "" || input.LastName == "" {
        return nil, errors.New("name or email is empty")
    }

    user := &users.User{
        FirstName:  input.FirstName,
        LastName:  input.LastName,
        Email: input.Email,
    }

    err := s.repo.Insert(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to insert user: %w", err)
    }

    return user, nil
}

