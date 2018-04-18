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
	NewSiteEndpoint             endpoint.Endpoint
	DeleteSiteEndpoint          endpoint.Endpoint
	WritePostEndpoint           endpoint.Endpoint
	RemovePostEndpoint          endpoint.Endpoint
	ReadPostEndpoint            endpoint.Endpoint
	WriteConfigEndpoint         endpoint.Endpoint
	ReadConfigEndpoint          endpoint.Endpoint
	UpdateAboutEndpoint         endpoint.Endpoint
	ReadAboutEndpoint           endpoint.Endpoint
	ChangeDefaultConfigEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger log.Logger) Set {
	return Set{
		NewSiteEndpoint:             Middlewares(MakeNewSiteEndpoint(svc), logger),
		DeleteSiteEndpoint:          Middlewares(MakeDeleteSiteEndpoint(svc), logger),
		WritePostEndpoint:           Middlewares(MakeWritePostEndpoint(svc), logger),
		RemovePostEndpoint:          Middlewares(MakeRemovePostEndpoint(svc), logger),
		ReadPostEndpoint:            Middlewares(MakeReadPostEndpoint(svc), logger),
		WriteConfigEndpoint:         Middlewares(MakeWriteConfigEndpoint(svc), logger),
		ReadConfigEndpoint:          Middlewares(MakeReadConfigEndpoint(svc), logger),
		UpdateAboutEndpoint:         Middlewares(MakeUpdateAboutEndpoint(svc), logger),
		ReadAboutEndpoint:           Middlewares(MakeReadAboutEndpoint(svc), logger),
		ChangeDefaultConfigEndpoint: Middlewares(MakeChangeDefaultConfigEndpoint(svc), logger),
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

// WriteConfig implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) WriteConfig(ctx context.Context, siteID uint, config string) error {
	resp, err := s.WriteConfigEndpoint(ctx, WriteConfigRequest{
		SiteID: siteID,
		Config: config,
	})
	if err != nil {
		return err
	}
	response := resp.(WriteConfigResponse)
	return response.Err
}

// ReadConfig implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) ReadConfig(ctx context.Context, siteID uint) (string, error) {
	resp, err := s.ReadConfigEndpoint(ctx, ReadConfigRequest{
		SiteID: siteID,
	})
	if err != nil {
		return "", err
	}
	response := resp.(ReadConfigResponse)
	return response.Config, response.Err
}

// UpdateAbout implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) UpdateAbout(ctx context.Context, siteID uint, content string) error {
	resp, err := s.UpdateAboutEndpoint(ctx, UpdateAboutRequest{
		SiteID:  siteID,
		Content: content,
	})
	if err != nil {
		return err
	}
	response := resp.(UpdateAboutResponse)
	return response.Err
}

// ReadAbout implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) ReadAbout(ctx context.Context, siteID uint) (string, error) {
	resp, err := s.ReadAboutEndpoint(ctx, ReadAboutRequest{
		SiteID: siteID,
	})
	if err != nil {
		return "", err
	}
	response := resp.(ReadAboutResponse)
	return response.Content, response.Err
}

// ChangeDefaultConfig implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) ChangeDefaultConfig(ctx context.Context, siteID uint, theme string) error {
	resp, err := s.ChangeDefaultConfigEndpoint(ctx, ChangeDefaultConfigRequest{
		SiteID: siteID,
		Theme:  theme,
	})
	if err != nil {
		return err
	}
	response := resp.(ChangeDefaultConfigResponse)
	return response.Err
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

// MakeWriteConfigEndpoint constructs a WriteConfig endpoint wrapping the service.
func MakeWriteConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WriteConfigRequest)
		err = svc.WriteConfig(ctx, req.SiteID, req.Config)
		return WriteConfigResponse{Err: err}, nil
	}
}

// MakeReadConfigEndpoint constructs a ReadConfig endpoint wrapping the service.
func MakeReadConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ReadConfigRequest)
		config, err := svc.ReadConfig(ctx, req.SiteID)
		return ReadConfigResponse{Config: config, Err: err}, nil
	}
}

// MakeUpdateAboutEndpoint constructs a UpdateAbout endpoint wrapping the service.
func MakeUpdateAboutEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateAboutRequest)
		err = svc.UpdateAbout(ctx, req.SiteID, req.Content)
		return UpdateAboutResponse{Err: err}, nil
	}
}

// MakeReadAboutEndpoint constructs a ReadAbout endpoint wrapping the service.
func MakeReadAboutEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ReadAboutRequest)
		content, err := svc.ReadAbout(ctx, req.SiteID)
		return ReadAboutResponse{Content: content, Err: err}, nil
	}
}

// MakeChangeDefaultConfigEndpoint constructs a ChangeDefaultConfig endpoint wrapping the service.
func MakeChangeDefaultConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ChangeDefaultConfigRequest)
		err = svc.ChangeDefaultConfig(ctx, req.SiteID, req.Theme)
		return ChangeDefaultConfigResponse{Err: err}, nil
	}
}
