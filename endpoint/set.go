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
}

// New returns a Set that wraps the provided server.
func New(svc service.Service, logger log.Logger) Set {
	return Set{
		NewSiteEndpoint:             Middlewares(MakeNewSiteEndpoint(svc), logger),
		DeleteSiteEndpoint:          Middlewares(MakeDeleteSiteEndpoint(svc), logger),
		CheckSitenameExistsEndpoint: Middlewares(MakeCheckSitenameExistsEndpoint(svc), logger),
		GetSiteIDByUserIDEndpoint:   Middlewares(MakeGetSiteIDByUserIDEndpoint(svc), logger),
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

// MakeNewSiteEndpoint constructs a NewSite endpoint wrapping the service.
func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		id, err := svc.NewSite(ctx, req.SiteName)
		return NewSiteResponse{SiteID: id, Err: err}, err
	}
}

// MakeDeleteSiteEndpoint constructs a DeleteSite endpoint wrapping the service.
func MakeDeleteSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteSiteRequest)
		err = svc.DeleteSite(ctx, req.SiteID)
		return DeleteSiteResponse{Err: err}, err
	}
}

// MakeCheckSitenameExistsEndpoint constructs a CheckSitenameExists endpoint wrapping the service.
func MakeCheckSitenameExistsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CheckSitenameExistsRequest)
		exists, err := svc.CheckSitenameExists(ctx, req.Sitename)
		return CheckSitenameExistsResponse{Exists: exists, Err: err}, err
	}
}

// MakeGetSiteIDByUserIDEndpoint constructs a GetSiteIDByUserID endpoint wrapping the service.
func MakeGetSiteIDByUserIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		siteID, err := svc.GetSiteIDByUserID(ctx)
		return GetSiteIDByUserIDResponse{SiteID: siteID, Err: err}, err
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
