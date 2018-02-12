package service

import (
	"context"
	"github.com/go-kit/kit/log"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
)

type Service interface {
	NewSite(ctx context.Context, email, sitename string) error
	DeleteSite(ctx context.Context, email, sitename string) error
}

func New(repoctlsvc repoctlservice.Service, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(repoctlsvc)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
}

func NewBasicService(repoctlsvc repoctlservice.Service) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
	}
}

func (s basicService) NewSite(ctx context.Context, email, sitename string) error {
	err := s.repoctlsvc.NewSite(ctx, email, sitename)
	if err != nil {
		return err
	}
	return nil
}

func (s basicService) DeleteSite(ctx context.Context, email, sitename string) error {
	return s.repoctlsvc.DeleteSite(ctx, email, sitename)
}
