package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/service"
)

type Set struct {
	NewSiteEndpoint             endpoint.Endpoint
	DeleteSiteEndpoint          endpoint.Endpoint
	CheckSitenameExistsEndpoint endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger) Set {
	var newsiteEndpoint endpoint.Endpoint
	{
		newsiteEndpoint = MakeNewSiteEndpoint(svc)
		newsiteEndpoint = LoggingMiddleware(logger)(newsiteEndpoint)
	}
	var deletesiteEndpoint endpoint.Endpoint
	{
		deletesiteEndpoint = MakeDeleteSiteEndpoint(svc)
		deletesiteEndpoint = LoggingMiddleware(logger)(deletesiteEndpoint)
	}
	var checkSitenameExistsEndpoint endpoint.Endpoint
	{
		checkSitenameExistsEndpoint = MakeCheckSitenameExistsEndpoint(svc)
		checkSitenameExistsEndpoint = LoggingMiddleware(logger)(checkSitenameExistsEndpoint)
	}
	return Set{
		NewSiteEndpoint:             newsiteEndpoint,
		DeleteSiteEndpoint:          deletesiteEndpoint,
		CheckSitenameExistsEndpoint: checkSitenameExistsEndpoint,
	}
}

func (s Set) NewSite(ctx context.Context, userID uint, sitename string) (uint, error) {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{UserId: userID, SiteName: sitename})
	if err != nil {
		return 0, err
	}
	response := resp.(NewSiteResponse)
	return response.SiteId, response.Err
}

func (s Set) DeleteSite(ctx context.Context, siteID uint) error {
	resp, err := s.DeleteSiteEndpoint(ctx, DeleteSiteRequest{SiteId: siteID})
	if err != nil {
		return err
	}
	response := resp.(DeleteSiteResponse)
	return response.Err
}

func (s Set) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	resp, err := s.CheckSitenameExistsEndpoint(ctx, CheckSitenameExistsRequest{Sitename: sitename})
	if err != nil {
		return true, err
	}
	response := resp.(CheckSitenameExistsResponse)
	return response.Exists, response.Err
}

func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		id, err := svc.NewSite(ctx, req.UserId, req.SiteName)
		return NewSiteResponse{SiteId: id, Err: err}, err
	}
}

func MakeDeleteSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteSiteRequest)
		err = svc.DeleteSite(ctx, req.SiteId)
		return DeleteSiteResponse{Err: err}, err
	}
}

func MakeCheckSitenameExistsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CheckSitenameExistsRequest)
		exists, err := svc.CheckSitenameExists(ctx, req.Sitename)
		return CheckSitenameExistsResponse{Exists: exists, Err: err}, err
	}
}

type Failer interface {
	Failed() error
}

func (r NewSiteResponse) Failed() error             { return r.Err }
func (r DeleteSiteResponse) Failed() error          { return r.Err }
func (r CheckSitenameExistsResponse) Failed() error { return r.Err }
