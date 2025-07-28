package connection

import (
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"project01/app/internal/db/config"
	"go.uber.org/zap"
	"fmt"
	"time"
	"context"
)

func ConnectionStart(ctx context.Context, cfg config.DBConfig, logger *zap.Logger) *pgxpool.Pool {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name,
	)
	poolConfig,err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Error("Error parsing pgx config:", zap.Error(err))
        log.Fatalf("failed to parse pgx config: %v", err)
    }

	poolConfig.MaxConns = 10
    poolConfig.MinConns = 1
    poolConfig.MaxConnIdleTime = time.Minute

	dbpool, err := pgxpool.NewWithConfig(ctx, poolConfig)
    if err != nil {
		logger.Error("Error creating pgx pool:", zap.Error(err))
        log.Fatalf("unable to create connection pool: %v", err)
    }

    if err := dbpool.Ping(ctx); err != nil {
		logger.Error("Error pinging database:", zap.Error(err))
        log.Fatalf("unable to ping database: %v", err)
    }

	logger.Info("Connected to PostgreSQL", zap.Time("time", time.Now()))

    return dbpool
}
