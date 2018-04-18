package service

//go:generate mockgen -destination=./mock/mock_service.go -package=mock github.com/seagullbird/headr-sitemgr/service Service

import (
	"context"
	"encoding/json"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-errors/errors"
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
	GetThemes(ctx context.Context, siteID uint) (string, error)
	UpdateSiteTheme(ctx context.Context, siteID uint, theme string) error
	PostAbout(ctx context.Context, siteID uint, content string) error
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

var (
	// ErrSiteNotFound indicates a site cannot be found
	ErrSiteNotFound = errors.New("site not found")
)

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
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return ErrSiteNotFound
	}

	// delete server service
	var delsiteEvent = mq.SiteUpdatedEvent{
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	if err := s.dispatcher.DispatchMessage("del_site_server", delsiteEvent); err != nil {
		return err
	}

	// delete database item
	if err := s.store.DeleteSite(site); err != nil {
		return err
	}
	// delete static files
	return s.repoctlsvc.DeleteSite(ctx, siteID)
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
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return "", ErrSiteNotFound
	}

	return s.repoctlsvc.ReadConfig(ctx, siteID)
}

func (s basicService) UpdateConfig(ctx context.Context, siteID uint, config string) error {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return ErrSiteNotFound
	}

	return s.repoctlsvc.WriteConfig(ctx, siteID, config)
}

func (s basicService) GetThemes(ctx context.Context, siteID uint) (string, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return "", ErrSiteNotFound
	}

	var (
		themes = []db.Theme{
			{
				Name:      "hugo-theme-cactus-plus",
				ThumbNail: "https://d33wubrfki0l68.cloudfront.net/a765cc66e8105ff5527be210beeba91d1e561902/49ef9/hugo-theme-cactus-plus/tn-featured-hugo-theme-cactus-plus_hu5badda48e9249d0a91a08133ff38d1ab_2179148_768x512_fill_catmullrom_top_2.png",
				HomePage:  "https://themes.gohugo.io/hugo-theme-cactus-plus/",
			},
			{
				Name:      "hyde-hyde",
				ThumbNail: "https://d33wubrfki0l68.cloudfront.net/03edb0da66032d8d5831bebace3cc2c31af0a2f1/fdcbf/hyde-hyde/tn-featured-hyde-hyde_hu6fa6cbe8ad4d9c1dab6f62cb035f7b86_209417_768x512_fill_catmullrom_top_2.png",
				HomePage:  "https://themes.gohugo.io/hyde-hyde/",
			},
		}
	)

	response := struct {
		Themes       []db.Theme `json:"themes"`
		CurrentTheme string     `json:"current_theme"`
	}{
		Themes:       themes,
		CurrentTheme: site.Theme,
	}

	respRaw, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(respRaw), nil
}

func (s basicService) UpdateSiteTheme(ctx context.Context, siteID uint, theme string) error {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return ErrSiteNotFound
	}

	site.Theme = theme
	err := s.store.PatchSite(site)
	if err != nil {
		return err
	}

	// https://stackoverflow.com/questions/11066946/partly-json-unmarshal-into-a-map-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
	originConfig, err := s.repoctlsvc.ReadConfig(ctx, siteID)
	if err != nil {
		return err
	}

	var configMap map[string]*json.RawMessage
	err = json.Unmarshal([]byte(originConfig), &configMap)
	if err != nil {
		return err
	}
	themeRaw, err := json.Marshal(theme)
	if err != nil {
		return err
	}
	*configMap["theme"] = themeRaw
	newConfigRaw, err := json.Marshal(configMap)

	return s.repoctlsvc.WriteConfig(ctx, siteID, string(newConfigRaw))
}

func (s basicService) PostAbout(ctx context.Context, siteID uint, content string) error {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	site, _ := s.store.GetSite(siteID)
	if site.UserID != userID {
		return ErrSiteNotFound
	}

	return s.repoctlsvc.UpdateAbout(ctx, siteID, content)
}
