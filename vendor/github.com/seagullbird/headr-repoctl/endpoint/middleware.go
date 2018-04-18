package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// Middlewares chains all middlewares together, returning the final endpoint.
// This is just a convenient method that helps in clearing up codes in endpoints.New
func Middlewares(e endpoint.Endpoint, logger log.Logger) endpoint.Endpoint {
	chain := endpoint.Chain(LoggingMiddleware(logger))
	return chain(e)
}
