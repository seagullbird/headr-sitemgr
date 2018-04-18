package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/service"
)

// Set collects all of the endpoints that compose an sitemgr service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	NewSiteEndpoint             endpoint.Endpoint
	DeleteSiteEndpoint          endpoint.Endpoint
	CheckSitenameExistsEndpoint endpoint.Endpoint
	GetSiteIDByUserIDEndpoint   endpoint.Endpoint
	GetConfigEndpoint           endpoint.Endpoint
	UpdateConfigEndpoint        endpoint.Endpoint
	GetThemesEndpoint           endpoint.Endpoint
	UpdateSiteThemeEndpoint     endpoint.Endpoint
	PostAboutEndpoint           endpoint.Endpoint
	GetAboutEndpoint            endpoint.Endpoint
}

// New returns a Set that wraps the provided server.
func New(svc service.Service, logger log.Logger) Set {
	return Set{
		NewSiteEndpoint:             Middlewares(MakeNewSiteEndpoint(svc), logger),
		DeleteSiteEndpoint:          Middlewares(MakeDeleteSiteEndpoint(svc), logger),
		CheckSitenameExistsEndpoint: Middlewares(MakeCheckSitenameExistsEndpoint(svc), logger),
		GetSiteIDByUserIDEndpoint:   Middlewares(MakeGetSiteIDByUserIDEndpoint(svc), logger),
		GetConfigEndpoint:           Middlewares(MakeGetConfigEndpoint(svc), logger),
		UpdateConfigEndpoint:        Middlewares(MakeUpdateConfigEndpoint(svc), logger),
		GetThemesEndpoint:           Middlewares(MakeGetThemesEndpoint(svc), logger),
		UpdateSiteThemeEndpoint:     Middlewares(MakeUpdateSiteThemeEndpoint(svc), logger),
		PostAboutEndpoint:           Middlewares(MakePostAboutEndpoint(svc), logger),
		GetAboutEndpoint:            Middlewares(MakeGetAboutEndpoint(svc), logger),
	}
}

// NewSite implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) NewSite(ctx context.Context, sitename string) (uint, error) {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{SiteName: sitename})
	if err != nil {
		return 0, err
	}
	response := resp.(NewSiteResponse)
	return response.SiteID, response.Err
}

// DeleteSite implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) DeleteSite(ctx context.Context, siteID uint) error {
	resp, err := s.DeleteSiteEndpoint(ctx, DeleteSiteRequest{SiteID: siteID})
	if err != nil {
		return err
	}
	response := resp.(DeleteSiteResponse)
	return response.Err
}

// CheckSitenameExists implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	resp, err := s.CheckSitenameExistsEndpoint(ctx, CheckSitenameExistsRequest{Sitename: sitename})
	if err != nil {
		return true, err
	}
	response := resp.(CheckSitenameExistsResponse)
	return response.Exists, response.Err
}

// GetSiteIDByUserID implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetSiteIDByUserID(ctx context.Context) (uint, error) {
	resp, err := s.GetSiteIDByUserIDEndpoint(ctx, GetSiteIDByUserIDRequest{})
	if err != nil {
		return 0, err
	}
	response := resp.(GetSiteIDByUserIDResponse)
	return response.SiteID, response.Err
}

// GetConfig implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetConfig(ctx context.Context, siteID uint) (string, error) {
	resp, err := s.GetConfigEndpoint(ctx, GetConfigRequest{SiteID: siteID})
	if err != nil {
		return "", err
	}
	response := resp.(GetConfigResponse)
	return response.Config, response.Err
}

// UpdateConfig implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) UpdateConfig(ctx context.Context, siteID uint, config string) error {
	resp, err := s.UpdateConfigEndpoint(ctx, UpdateConfigRequest{SiteID: siteID, Config: config})
	if err != nil {
		return err
	}
	response := resp.(UpdateConfigResponse)
	return response.Err
}

// GetThemes implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetThemes(ctx context.Context, siteID uint) (string, error) {
	resp, err := s.GetThemesEndpoint(ctx, GetThemesRequest{SiteID: siteID})
	if err != nil {
		return "", err
	}
	response := resp.(GetThemesResponse)
	return response.Themes, response.Err
}

// UpdateSiteTheme implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) UpdateSiteTheme(ctx context.Context, siteID uint, theme string) error {
	resp, err := s.UpdateSiteThemeEndpoint(ctx, UpdateSiteThemeRequest{SiteID: siteID, Theme: theme})
	if err != nil {
		return err
	}
	response := resp.(UpdateSiteThemeResponse)
	return response.Err
}

// PostAbout implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) PostAbout(ctx context.Context, siteID uint, content string) error {
	resp, err := s.PostAboutEndpoint(ctx, PostAboutRequest{SiteID: siteID, Content: content})
	if err != nil {
		return err
	}
	response := resp.(PostAboutResponse)
	return response.Err
}

// GetAbout implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetAbout(ctx context.Context, siteID uint) (string, error) {
	resp, err := s.GetAboutEndpoint(ctx, GetAboutRequest{SiteID: siteID})
	if err != nil {
		return "", err
	}
	response := resp.(GetAboutResponse)
	return response.Content, response.Err
}

// MakeNewSiteEndpoint constructs a NewSite endpoint wrapping the service.
func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		id, err := svc.NewSite(ctx, req.SiteName)
		return NewSiteResponse{SiteID: id, Err: err}, nil
	}
}

// MakeDeleteSiteEndpoint constructs a DeleteSite endpoint wrapping the service.
func MakeDeleteSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteSiteRequest)
		err = svc.DeleteSite(ctx, req.SiteID)
		return DeleteSiteResponse{Err: err}, nil
	}
}

// MakeCheckSitenameExistsEndpoint constructs a CheckSitenameExists endpoint wrapping the service.
func MakeCheckSitenameExistsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CheckSitenameExistsRequest)
		exists, err := svc.CheckSitenameExists(ctx, req.Sitename)
		return CheckSitenameExistsResponse{Exists: exists, Err: err}, nil
	}
}

// MakeGetSiteIDByUserIDEndpoint constructs a GetSiteIDByUserID endpoint wrapping the service.
func MakeGetSiteIDByUserIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		siteID, err := svc.GetSiteIDByUserID(ctx)
		return GetSiteIDByUserIDResponse{SiteID: siteID, Err: err}, nil
	}
}

// MakeGetConfigEndpoint constructs a GetConfig endpoint wrapping the service.
func MakeGetConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetConfigRequest)
		config, err := svc.GetConfig(ctx, req.SiteID)
		return GetConfigResponse{Config: config, Err: err}, nil
	}
}

// MakeUpdateConfigEndpoint constructs a UpdateConfig endpoint wrapping the service.
func MakeUpdateConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateConfigRequest)
		err = svc.UpdateConfig(ctx, req.SiteID, req.Config)
		return UpdateConfigResponse{Err: err}, nil
	}
}

// MakeGetThemesEndpoint constructs a GetThemes endpoint wrapping the service.
func MakeGetThemesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetThemesRequest)
		themes, err := svc.GetThemes(ctx, req.SiteID)
		return GetThemesResponse{Themes: themes, Err: err}, nil
	}
}

// MakeUpdateSiteThemeEndpoint constructs a UpdateSiteTheme endpoint wrapping the service.
func MakeUpdateSiteThemeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateSiteThemeRequest)
		err = svc.UpdateSiteTheme(ctx, req.SiteID, req.Theme)
		return UpdateSiteThemeResponse{Err: err}, nil
	}
}

// MakePostAboutEndpoint constructs a PostAbout endpoint wrapping the service.
func MakePostAboutEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostAboutRequest)
		err = svc.PostAbout(ctx, req.SiteID, req.Content)
		return PostAboutResponse{Err: err}, nil
	}
}

// MakeGetAboutEndpoint constructs a GetAbout endpoint wrapping the service.
func MakeGetAboutEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAboutRequest)
		content, err := svc.GetAbout(ctx, req.SiteID)
		return GetAboutResponse{Content: content, Err: err}, nil
	}
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so if they've
// failed, and if so encode them using a separate write path based on the error.
type Failer interface {
	Failed() error
}

// Failed implements Failer.
func (r NewSiteResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r DeleteSiteResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r CheckSitenameExistsResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetSiteIDByUserIDResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetConfigResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r UpdateConfigResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetThemesResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r UpdateSiteThemeResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r PostAboutResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetAboutResponse) Failed() error { return r.Err }
