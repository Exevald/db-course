package common

import (
	"context"
	"net/http"
	"time"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	log "github.com/sirupsen/logrus"

	"orgchart/pkg/orgchart/common/logging"
)

func NewLoggingMiddleware(logger *log.Logger) strictnethttp.StrictHTTPMiddlewareFunc {
	return func(f strictnethttp.StrictHTTPHandlerFunc, operationID string) strictnethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
			start := time.Now()
			resp, err := f(ctx, w, r, request)
			duration := time.Since(start).String()
			fields := log.Fields{
				"args":     logging.TrimForLogs(request, logging.DefaultTrimForLogsOpts),
				"duration": duration,
				"method":   operationID,
			}
			loggerWithFields := logger.WithFields(fields)
			if err != nil {
				loggerWithFields.Error(err, "call failed")
			} else {
				loggerWithFields.Info("call succeeded")
			}
			return resp, err
		}
	}
}
