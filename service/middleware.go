package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// Logging Middleware
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			logger,
			next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) NewSite(ctx context.Context, userID uint, sitename string) (err error) {
	err = mw.next.NewSite(ctx, userID, sitename)
	mw.logger.Log("method", "NewSite", "userID", userID, "sitename", sitename, "err", err)
	return
}

func (mw loggingMiddleware) DeleteSite(ctx context.Context, siteID uint) (err error) {
	err = mw.next.DeleteSite(ctx, siteID)
	mw.logger.Log("method", "DeleteSite", "siteID", siteID, "err", err)
	return
}
