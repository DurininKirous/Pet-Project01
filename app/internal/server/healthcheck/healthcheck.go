package healthcheckRouter

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"project01/app/internal/handlers/healthcheck"
	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger) http.Handler {
	r := chi.NewRouter()
	r.Get("/", healthcheckService.PingHandler(logger))
	return r
}

