package service

//go:generate mockgen -destination=./mock/mock_service.go -package=mock github.com/seagullbird/headr-sitemgr/service Service

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
	GetConfig(ctx context.Context, siteID uint) (string, error)
	UpdateConfig(ctx context.Context, siteID uint, config string) error
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
	err = s.repoctlsvc.NewSite(ctx, siteID, site.Theme)
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
	if err := s.store.DeleteSite(site); err != nil {
		return err
	}
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
	exists, err := s.store.CheckSitenameExists(sitename)
	if err != nil {
		return false, nil
	}
	return exists, nil
}

func (s basicService) GetSiteIDByUserID(ctx context.Context) (uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"]
	siteID, _ := s.store.GetSiteIDByUserID(userID.(string))
	return siteID, nil
}

func (s basicService) GetConfig(ctx context.Context, siteID uint) (string, error) {
	return s.repoctlsvc.ReadConfig(ctx, siteID)
}

func (s basicService) UpdateConfig(ctx context.Context, siteID uint, config string) error {
	return s.repoctlsvc.WriteConfig(ctx, siteID, config)
}
