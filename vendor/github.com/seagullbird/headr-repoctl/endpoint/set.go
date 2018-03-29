package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/service"
)

// Set collects all of the endpoints that compose an repoctl service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	NewSiteEndpoint    endpoint.Endpoint
	DeleteSiteEndpoint endpoint.Endpoint
	WritePostEndpoint  endpoint.Endpoint
	RemovePostEndpoint endpoint.Endpoint
	ReadPostEndpoint   endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger log.Logger) Set {
	return Set{
		NewSiteEndpoint:    Middlewares(MakeNewSiteEndpoint(svc), logger),
		DeleteSiteEndpoint: Middlewares(MakeDeleteSiteEndpoint(svc), logger),
		WritePostEndpoint:  Middlewares(MakeWritePostEndpoint(svc), logger),
		RemovePostEndpoint: Middlewares(MakeRemovePostEndpoint(svc), logger),
		ReadPostEndpoint:   Middlewares(MakeReadPostEndpoint(svc), logger),
	}
}

// NewSite implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) NewSite(ctx context.Context, siteID uint, theme string) error {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{SiteID: siteID, Theme: theme})
	if err != nil {
		return err
	}
	response := resp.(NewSiteResponse)
	return response.Err
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

// WritePost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) WritePost(ctx context.Context, siteID uint, filename, content string) error {
	resp, err := s.WritePostEndpoint(ctx, WritePostRequest{
		SiteID:   siteID,
		Filename: filename,
		Content:  content,
	})
	if err != nil {
		return err
	}
	response := resp.(WritePostResponse)
	return response.Err
}

// RemovePost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) RemovePost(ctx context.Context, siteID uint, filename string) error {
	resp, err := s.RemovePostEndpoint(ctx, RemovePostRequest{
		SiteID:   siteID,
		Filename: filename,
	})
	if err != nil {
		return err
	}
	response := resp.(RemovePostResponse)
	return response.Err
}

// ReadPost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) ReadPost(ctx context.Context, siteID uint, filename string) (string, error) {
	resp, err := s.ReadPostEndpoint(ctx, ReadPostRequest{
		SiteID:   siteID,
		Filename: filename,
	})
	if err != nil {
		return "", err
	}
	response := resp.(ReadPostResponse)
	return response.Content, response.Err
}

// MakeNewSiteEndpoint constructs a NewSite endpoint wrapping the service.
func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		err = svc.NewSite(ctx, req.SiteID, req.Theme)
		return NewSiteResponse{Err: err}, nil
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

// MakeWritePostEndpoint constructs a WritePost endpoint wrapping the service.
func MakeWritePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WritePostRequest)
		err = svc.WritePost(ctx, req.SiteID, req.Filename, req.Content)
		return WritePostResponse{Err: err}, nil
	}
}

// MakeRemovePostEndpoint constructs a RemovePost endpoint wrapping the service.
func MakeRemovePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RemovePostRequest)
		err = svc.RemovePost(ctx, req.SiteID, req.Filename)
		return RemovePostResponse{Err: err}, nil
	}
}

// MakeReadPostEndpoint constructs a ReadPost endpoint wrapping the service.
func MakeReadPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ReadPostRequest)
		content, err := svc.ReadPost(ctx, req.SiteID, req.Filename)
		return ReadPostResponse{Content: content, Err: err}, nil
	}
}
