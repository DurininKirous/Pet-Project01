package usersService

import (
	users "project01/app/internal/db/users"
	"context"
	"fmt"
	"errors"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"strconv"
)

type Service struct {
	repo *users.Repo
	logger *zap.Logger
}

type CreateUserInput struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
}

func New(repo *users.Repo, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetAll(ctx context.Context) ([]users.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
		   s.logger.Info("Getting users from table users: user not found", zap.Error(err))
           return nil, fmt.Errorf("user not found: %w", err)
        }
		s.logger.Error("Error getting users from table users", zap.Error(err))
		return nil, fmt.Errorf("service: failed to get users: %w", err)
	}
	s.logger.Info("Got users from table users", zap.Error(err))
	return users, nil
}

func (s *Service) GetById(ctx context.Context, id int) (*users.User, error) {
		
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
		   s.logger.Info("Getting users from table users: user not found", zap.Error(err))
           return nil, fmt.Errorf("user not found: %w", err)
        }
		s.logger.Error("Error getting users from table users", zap.Error(err))
		return nil, fmt.Errorf("service: failed to get user: %w", err)
	}
	s.logger.Info("Got user from table users", zap.String("id", strconv.Itoa(user.ID)), zap.Error(err))
	return user, nil
}

func (s *Service) Create(ctx context.Context, input CreateUserInput) (*users.User, error) {
    if input.FirstName == "" || input.Email == "" || input.LastName == "" {
		s.logger.Error("Error: name or email is empty", zap.String("env","dev"))
        return nil, errors.New("name or email is empty")
    }

    user := &users.User{
        FirstName:  input.FirstName,
        LastName:  input.LastName,
        Email: input.Email,
    }

    err := s.repo.Insert(ctx, user)
    if err != nil {
		s.logger.Error("Error inserting user into table users", zap.Error(err))
        return nil, fmt.Errorf("failed to insert user: %w", err)
    }
	s.logger.Info("Inserted user into table users", zap.String("id", strconv.Itoa(user.ID)))

    return user, nil
}

