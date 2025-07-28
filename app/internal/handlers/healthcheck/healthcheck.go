package healthcheckService

import (
	"net/http"
	"go.uber.org/zap"
)

func PingHandler(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Healthcheck ping received",
			zap.String("remote", r.RemoteAddr),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

