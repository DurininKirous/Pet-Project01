package connection

import (
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"project01/app/internal/db/config"
	"fmt"
	"time"
	"context"
)

func ConnectionStart(ctx context.Context, cfg config.DBConfig) *pgxpool.Pool {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name,
	)
	poolConfig,err := pgxpool.ParseConfig(connStr)
	if err != nil {
        log.Fatalf("failed to parse pgx config: %v", err)
    }

	poolConfig.MaxConns = 10
    poolConfig.MinConns = 1
    poolConfig.MaxConnIdleTime = time.Minute

	dbpool, err := pgxpool.NewWithConfig(ctx, poolConfig)
    if err != nil {
        log.Fatalf("unable to create connection pool: %v", err)
    }

	// тестируем подключение
    if err := dbpool.Ping(ctx); err != nil {
        log.Fatalf("unable to ping database: %v", err)
    }

    log.Println("Connected to PostgreSQL")

    return dbpool
}
