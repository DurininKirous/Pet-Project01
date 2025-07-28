package usersRouter

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	users "project01/app/internal/handlers/users"
	"project01/app/internal/service/users"
	db "project01/app/internal/db/users"
	"go.uber.org/zap"
)

func NewRouter(repo *db.Repo, logger *zap.Logger) http.Handler {
    r := chi.NewRouter()
	service := usersService.New(repo, logger)
    handler := users.New(service) 

    r.Get("/", handler.GetAll)
    r.Get("/{id}", handler.GetByID)
    r.Post("/", handler.Create)

    return r
}
