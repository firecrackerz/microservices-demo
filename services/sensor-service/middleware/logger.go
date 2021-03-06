package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/moorara/microservices-demo/services/sensor-service/util"
)

type loggerMiddleware struct {
	logger log.Logger
}

// NewLoggerMiddleware creates a new middleware for logging
func NewLoggerMiddleware(logger log.Logger) Middleware {
	return &loggerMiddleware{
		logger: logger,
	}
}

func (m *loggerMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		endpoint := r.URL.Path
		headers := r.Header

		// This only works with mux router
		for p, v := range mux.Vars(r) {
			endpoint = strings.Replace(endpoint, v, ":"+p, 1)
		}

		rw := util.NewResponseWriter(w)
		next(rw, r)

		duration := time.Now().Sub(start).Seconds()
		durationMS := int(duration * 1000)
		statusCode := rw.StatusCode()
		statusClass := rw.StatusClass()

		result := []interface{}{
			"req.method", method,
			"req.endpoint", endpoint,
			"req.headers", headers,
			"res.statusCode", statusCode,
			"res.statusClass", statusClass,
			"responseTime", durationMS,
			"message", fmt.Sprintf("%s %s %d %d", method, endpoint, statusCode, durationMS),
		}

		switch {
		case statusCode >= 500:
			level.Error(m.logger).Log(result...)
		case statusCode >= 400:
			level.Warn(m.logger).Log(result...)
		case statusCode >= 100:
			level.Info(m.logger).Log(result...)
		}
	}
}
