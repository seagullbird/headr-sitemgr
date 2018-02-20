package service

import (
	"context"
	"github.com/go-kit/kit/log"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"time"
)

type Service interface {
	NewSite(ctx context.Context, email, sitename string) error
	DeleteSite(ctx context.Context, email, sitename string) error
}

func New(repoctlsvc repoctlservice.Service, logger log.Logger, dispatcher dispatch.Dispatcher) Service {
	var svc Service
	{
		svc = NewBasicService(repoctlsvc, dispatcher)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	dispatcher dispatch.Dispatcher
}

func NewBasicService(repoctlsvc repoctlservice.Service, dispatcher dispatch.Dispatcher) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
		dispatcher: dispatcher,
	}
}

func (s basicService) NewSite(ctx context.Context, email, sitename string) error {
	err := s.repoctlsvc.NewSite(ctx, email, sitename)
	if err != nil {
		return err
	}
	var newsiteEvent = mq.NewSiteEvent{
		Email: email,
		SiteName: sitename,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("new_site_server", newsiteEvent)
}

func (s basicService) DeleteSite(ctx context.Context, email, sitename string) error {
	err := s.repoctlsvc.DeleteSite(ctx, email, sitename)
	if err != nil {
		return err
	}
	var delsiteEvent = mq.DelSiteEvent{
		Email: email,
		SiteName: sitename,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("del_site_server", delsiteEvent)
}
