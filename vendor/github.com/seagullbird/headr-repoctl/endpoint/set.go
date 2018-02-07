package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/go-kit/kit/log"
)

type Set struct {
	NewSiteEndpoint		endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger) Set {
	var newsiteEndpoint endpoint.Endpoint
	{
		newsiteEndpoint = MakeNewSiteEndpoint(svc)
		newsiteEndpoint = LoggingMiddleware(logger)(newsiteEndpoint)
	}
	return Set{
		NewSiteEndpoint: newsiteEndpoint,
	}
}

func (s Set) NewSite(ctx context.Context, email, sitename string) error {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{Email: email, SiteName: sitename})
	if err != nil {
		return err
	}
	response := resp.(NewSiteResponse)
	return response.Err
}

func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		err = svc.NewSite(ctx, req.Email, req.SiteName)
		return NewSiteResponse{Err: err}, nil
	}
}
