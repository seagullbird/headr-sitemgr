package service

import (
	"context"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
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
	NewSite(ctx context.Context, sitename string) (uint, error)
	DeleteSite(ctx context.Context, siteID uint) error
	CheckSitenameExists(ctx context.Context, sitename string) (bool, error)
	GetSiteIDByUserID(ctx context.Context) (uint, error)
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

func (s basicService) NewSite(ctx context.Context, sitename string) (uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
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
		SiteID:     site.Model.ID,
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
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("del_site_server", delsiteEvent)
}

func (s basicService) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	return s.store.CheckSitenameExists(sitename)
}

func (s basicService) GetSiteIDByUserID(ctx context.Context) (uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"]
	return s.store.GetSiteIDByUserID(userID.(string))
}

// EmptyService is only used for transport tests
type EmptyService struct{}

// NewSite implements Service.NewSite
func (e EmptyService) NewSite(ctx context.Context, sitename string) (uint, error) {
	return 0, nil
}

// DeleteSite implements Service.DeleteSite
func (e EmptyService) DeleteSite(ctx context.Context, siteID uint) error {
	return nil
}

// CheckSitenameExists implements Service.CheckSitenameExists
func (e EmptyService) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	return true, nil
}

// GetSiteIDByUserID implements Service.GetSiteIDByUserID
func (e EmptyService) GetSiteIDByUserID(ctx context.Context) (uint, error) {
	return 0, nil
}
