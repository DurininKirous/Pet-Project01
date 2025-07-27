package healthcheckRouter

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"project01/app/internal/handlers/healthcheck"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", healthcheckService.PingHandler)
	return r
}
