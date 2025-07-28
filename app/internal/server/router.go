package server

import (
	"net/http"
	"github.com/go-chi/chi/v5"
    //"github.com/go-chi/chi/v5/middleware"

	"project01/app/internal/server/healthcheck"
	"project01/app/internal/server/users"
	db "project01/app/internal/db/users"
	"go.uber.org/zap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(repo *db.Repo, logger *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Mount("/ping", healthcheckRouter.NewRouter(logger))
    r.Mount("/users", usersRouter.NewRouter(repo, logger))
	r.Mount("/metrics", promhttp.Handler())

	return r
}

