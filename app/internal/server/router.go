package server

import (
	"net/http"
	"github.com/go-chi/chi/v5"
    //"github.com/go-chi/chi/v5/middleware"

	"project01/app/internal/server/healthcheck"
	"project01/app/internal/server/users"
	db "project01/app/internal/db/users"
)

func NewRouter(repo *db.Repo) http.Handler {
	r := chi.NewRouter()

	r.Mount("/ping", healthcheckRouter.NewRouter())
    r.Mount("/users", usersRouter.NewRouter(repo))
	//r.Mount("/logs", logsHandler.NewRouter())

	return r
}

