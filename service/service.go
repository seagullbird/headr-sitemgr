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

// Service describes a service that deals with site management operations (sitemgr).
type Service interface {
	NewSite(ctx context.Context, userID uint, sitename string) (uint, error)
	DeleteSite(ctx context.Context, siteID uint) error
	CheckSitenameExists(ctx context.Context, sitename string) (bool, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(repoctlsvc repoctlservice.Service, logger log.Logger, dispatcher dispatch.Dispatcher, store db.Store) Service {
	var svc Service
	{
		svc = newBasicService(repoctlsvc, dispatcher, store)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	dispatcher dispatch.Dispatcher
	store      db.Store
}

func newBasicService(repoctlsvc repoctlservice.Service, dispatcher dispatch.Dispatcher, store db.Store) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
		dispatcher: dispatcher,
		store:      store,
	}
}

func (s basicService) NewSite(ctx context.Context, userID uint, sitename string) (uint, error) {
	site := &db.Site{
		UserID:   userID,
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
	// delete database item
	site, _ := s.store.GetSite(siteID)
	s.store.DeleteSite(site)
	// delete static files
	err := s.repoctlsvc.DeleteSite(ctx, siteID)
	if err != nil {
		return err
	}
	// delete server service
	var delsiteEvent = mq.SiteUpdatedEvent{
		SiteId:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("del_site_server", delsiteEvent)
}

func (s basicService) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	return s.store.CheckSitenameExists(sitename)
}

// EmptyService is only used for transport tests
type EmptyService struct{}

// NewSite inplements Service.NewSite
func (e EmptyService) NewSite(ctx context.Context, userID uint, sitename string) (uint, error) {
	return 0, nil
}

// DeleteSite inplements Service.DeleteSite
func (e EmptyService) DeleteSite(ctx context.Context, siteID uint) error {
	return nil
}

// CheckSitenameExists inplements Service.CheckSitenameExists
func (e EmptyService) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	return true, nil
}
