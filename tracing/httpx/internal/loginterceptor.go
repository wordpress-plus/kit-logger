package internal

import (
	"github.com/wordpress-plus/kit-logger"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/propagation"
)

func LogInterceptor(r *http.Request) (*http.Request, ResponseHandler) {
	start := time.Now()
	return r, func(resp *http.Response, err error) {
		duration := time.Since(start)
		if err != nil {
			logger := kit.Logger.WithContext(r.Context()).WithDuration(duration)
			logger.Errorf("[HTTP] %s %s - %v", r.Method, r.URL, err)
			return
		}

		var tc propagation.TraceContext
		ctx := tc.Extract(r.Context(), propagation.HeaderCarrier(resp.Header))
		logger := kit.Logger.WithContext(ctx).WithDuration(duration)
		if isOkResponse(resp.StatusCode) {
			logger.Infof("[HTTP] %d - %s %s", resp.StatusCode, r.Method, r.URL)
		} else {
			logger.Errorf("[HTTP] %d - %s %s", resp.StatusCode, r.Method, r.URL)
		}
	}
}

func isOkResponse(code int) bool {
	return code < http.StatusBadRequest
}
