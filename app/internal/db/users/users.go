package usersDBMethods

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"time"
	"go.uber.org/zap"
)

type Repo struct {
	db *pgxpool.Pool
	logger *zap.Logger
}

func New(pool *pgxpool.Pool, logger *zap.Logger) *Repo {
	return &Repo{db: pool, logger: logger}
}

type User struct {
	ID 		   int 	    `json:"id"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	Email 	   string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Repo) FindAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, firstname, lastname, email, created_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("Error getting users from table users", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt); err != nil {
			r.logger.Error("Error scanning rows in user's data", zap.Error(err))
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repo) FindByID(ctx context.Context, id int) (*User, error) {
	query := `SELECT id, firstname, lastname, email, created_at FROM users WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		r.logger.Info("Error scanning rows in user's data", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (r *Repo) Insert(ctx context.Context, user *User) error {
	query := `INSERT INTO users (firstname, lastname, email) values ($1, $2, $3) RETURNING id, created_at`

	return r.db.QueryRow(ctx, query, user.FirstName, user.LastName, user.Email).Scan(&user.ID, &user.CreatedAt)
}
