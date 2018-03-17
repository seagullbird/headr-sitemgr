package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/service"
)

type Set struct {
	NewSiteEndpoint    endpoint.Endpoint
	DeleteSiteEndpoint endpoint.Endpoint
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
	return Set{
		NewSiteEndpoint:    newsiteEndpoint,
		DeleteSiteEndpoint: deletesiteEndpoint,
	}
}

func (s Set) NewSite(ctx context.Context, userID uint, sitename string) error {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{UserId: userID, SiteName: sitename})
	if err != nil {
		return err
	}
	response := resp.(NewSiteResponse)
	return response.Err
}

func (s Set) DeleteSite(ctx context.Context, siteID uint) error {
	resp, err := s.DeleteSiteEndpoint(ctx, DeleteSiteRequest{SiteId: siteID})
	if err != nil {
		return err
	}
	response := resp.(DeleteSiteResponse)
	return response.Err
}

func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		err = svc.NewSite(ctx, req.UserId, req.SiteName)
		return NewSiteResponse{Err: err}, err
	}
}

func MakeDeleteSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteSiteRequest)
		err = svc.DeleteSite(ctx, req.SiteId)
		return DeleteSiteResponse{Err: err}, err
	}
}

type Failer interface {
	Failed() error
}

func (r NewSiteResponse) Failed() error { return r.Err }
