package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/dispatch"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-sitemgr/config"
	"github.com/seagullbird/headr-sitemgr/db"
	"time"
)

type Service interface {
	NewSite(ctx context.Context, userID uint, sitename string) (uint, error)
	DeleteSite(ctx context.Context, siteID uint) error
	CheckSitenameExists(ctx context.Context, sitename string) (bool, error)
}

func New(repoctlsvc repoctlservice.Service, logger log.Logger, dispatcher dispatch.Dispatcher, store db.Store) Service {
	var svc Service
	{
		svc = NewBasicService(repoctlsvc, dispatcher, store)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	dispatcher dispatch.Dispatcher
	store      db.Store
}

func NewBasicService(repoctlsvc repoctlservice.Service, dispatcher dispatch.Dispatcher, store db.Store) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
		dispatcher: dispatcher,
		store:      store,
	}
}

// TODO: After creating new site, sitemgr should be responsible for updating siteID to user's app_metadata through Auth0's Management API
func (s basicService) NewSite(ctx context.Context, userID uint, sitename string) (uint, error) {
	site := &db.Site{
		UserId:   userID,
		Theme:    config.InitialTheme,
		Sitename: sitename,
	}
	siteID, err := s.store.InsertSite(site)
	if err != nil {
		return 0, err
	}
	err = s.repoctlsvc.NewSite(ctx, siteID)
	if err != nil {
		return 0, err
	}
	var newsiteEvent = mq.SiteUpdatedEvent{
		SiteId:     site.Model.ID,
		Theme:      site.Theme,
		ReceivedOn: time.Now().Unix(),
	}
	return site.Model.ID, s.dispatcher.DispatchMessage("new_site_server", newsiteEvent)
}

func (s basicService) DeleteSite(ctx context.Context, siteID uint) error {
	err := s.repoctlsvc.DeleteSite(ctx, siteID)
	if err != nil {
		return err
	}
	var delsiteEvent = mq.SiteUpdatedEvent{
		SiteId:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("del_site_server", delsiteEvent)
}

func (s basicService) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	return s.store.CheckSitenameExists(sitename)
}
