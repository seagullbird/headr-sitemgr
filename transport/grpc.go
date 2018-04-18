package transport

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/pb"
	"github.com/seagullbird/headr-sitemgr/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	newsite             grpctransport.Handler
	deletesite          grpctransport.Handler
	checksitenameexists grpctransport.Handler
	getsiteidbyuserid   grpctransport.Handler
	getconfig           grpctransport.Handler
	updateconfig        grpctransport.Handler
	getthemes           grpctransport.Handler
	updatesitetheme     grpctransport.Handler
	postabout           grpctransport.Handler
	getabout            grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC SitemgrServer.
func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.SitemgrServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	}
	return &grpcServer{
		newsite: grpctransport.NewServer(
			endpoints.NewSiteEndpoint,
			decodeGRPCNewSiteRequest,
			encodeGRPCNewSiteResponse,
			options...,
		),
		deletesite: grpctransport.NewServer(
			endpoints.DeleteSiteEndpoint,
			decodeGRPCDeleteSiteRequest,
			encodeGRPCDeleteSiteResponse,
			options...,
		),
		checksitenameexists: grpctransport.NewServer(
			endpoints.CheckSitenameExistsEndpoint,
			decodeGRPCCheckSitenameExistsRequest,
			encodeGRPCCheckSitenameExistsResponse,
			options...,
		),
		getsiteidbyuserid: grpctransport.NewServer(
			endpoints.GetSiteIDByUserIDEndpoint,
			decodeGRPCGetSiteIDByUserIDRequest,
			encodeGRPCGetSiteIDByUserIDResponse,
			options...,
		),
		getconfig: grpctransport.NewServer(
			endpoints.GetConfigEndpoint,
			decodeGRPCGetConfigRequest,
			encodeGRPCGetConfigResponse,
			options...,
		),
		updateconfig: grpctransport.NewServer(
			endpoints.UpdateConfigEndpoint,
			decodeGRPCUpdateConfigRequest,
			encodeGRPCUpdateConfigResponse,
			options...,
		),
		getthemes: grpctransport.NewServer(
			endpoints.GetThemesEndpoint,
			decodeGRPCGetThemesRequest,
			encodeGRPCGetThemesResponse,
			options...,
		),
		updatesitetheme: grpctransport.NewServer(
			endpoints.UpdateSiteThemeEndpoint,
			decodeGRPCUpdateSiteThemeRequest,
			encodeGRPCUpdateSiteThemeResponse,
			options...,
		),
		postabout: grpctransport.NewServer(
			endpoints.PostAboutEndpoint,
			decodeGRPCPostAboutRequest,
			encodeGRPCPostAboutResponse,
			options...,
		),
		getabout: grpctransport.NewServer(
			endpoints.GetAboutEndpoint,
			decodeGRPCGetAboutRequest,
			encodeGRPCGetAboutResponse,
			options...,
		),
	}
}

// NewGRPCClient returns an SitemgrService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport.
func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
	}
	var newsiteEndpoint kitendpoint.Endpoint
	{
		newsiteEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"NewSite",
			encodeGRPCNewSiteRequest,
			decodeGRPCNewSiteResponse,
			pb.CreateNewSiteReply{},
			options...,
		).Endpoint()
	}
	var deletesiteEndpoint kitendpoint.Endpoint
	{
		deletesiteEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"DeleteSite",
			encodeGRPCDeleteSiteRequest,
			decodeGRPCDeleteSiteResponse,
			pb.ProxyDeleteSiteReply{},
			options...,
		).Endpoint()
	}
	var checksitenameexistsEndpoint kitendpoint.Endpoint
	{
		checksitenameexistsEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"CheckSitenameExists",
			encodeGRPCCheckSitenameExistsRequest,
			decodeGRPCCheckSitenameExistsResponse,
			pb.CheckSitenameExistsReply{},
			options...,
		).Endpoint()
	}
	var getsiteidbyuseridEndpoint kitendpoint.Endpoint
	{
		getsiteidbyuseridEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"GetSiteIDByUserID",
			encodeGRPCGetSiteIDByUserIDRequest,
			decodeGRPCGetSiteIDByUserIDResponse,
			pb.GetSiteIDByUserIDReply{},
			options...,
		).Endpoint()
	}
	var getconfigEndpoint kitendpoint.Endpoint
	{
		getconfigEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"GetConfig",
			encodeGRPCGetConfigRequest,
			decodeGRPCGetConfigResponse,
			pb.GetConfigReply{},
			options...,
		).Endpoint()
	}
	var updateconfigEndpoint kitendpoint.Endpoint
	{
		updateconfigEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"UpdateConfig",
			encodeGRPCUpdateConfigRequest,
			decodeGRPCUpdateConfigResponse,
			pb.UpdateConfigReply{},
			options...,
		).Endpoint()
	}
	var getthemesEndpoint kitendpoint.Endpoint
	{
		getthemesEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"GetThemes",
			encodeGRPCGetThemesRequest,
			decodeGRPCGetThemesResponse,
			pb.GetThemesReply{},
			options...,
		).Endpoint()
	}
	var updatesitethemeEndpoint kitendpoint.Endpoint
	{
		updatesitethemeEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"UpdateSiteTheme",
			encodeGRPCUpdateSiteThemeRequest,
			decodeGRPCUpdateSiteThemeResponse,
			pb.UpdateSiteThemeReply{},
			options...,
		).Endpoint()
	}
	var postaboutEndpoint kitendpoint.Endpoint
	{
		postaboutEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"PostAbout",
			encodeGRPCPostAboutRequest,
			decodeGRPCPostAboutResponse,
			pb.PostAboutReply{},
			options...,
		).Endpoint()
	}
	var getaboutEndpoint kitendpoint.Endpoint
	{
		getaboutEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"GetAbout",
			encodeGRPCGetAboutRequest,
			decodeGRPCGetAboutResponse,
			pb.GetAboutReply{},
			options...,
		).Endpoint()
	}
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint:             newsiteEndpoint,
		DeleteSiteEndpoint:          deletesiteEndpoint,
		CheckSitenameExistsEndpoint: checksitenameexistsEndpoint,
		GetSiteIDByUserIDEndpoint:   getsiteidbyuseridEndpoint,
		GetConfigEndpoint:           getconfigEndpoint,
		UpdateConfigEndpoint:        updateconfigEndpoint,
		GetThemesEndpoint:           getthemesEndpoint,
		UpdateSiteThemeEndpoint:     updatesitethemeEndpoint,
		PostAboutEndpoint:           postaboutEndpoint,
		GetAboutEndpoint:            getaboutEndpoint,
	}
}

func (s *grpcServer) NewSite(ctx context.Context, req *pb.CreateNewSiteRequest) (*pb.CreateNewSiteReply, error) {
	_, rep, err := s.newsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateNewSiteReply), nil
}

func (s *grpcServer) DeleteSite(ctx context.Context, req *pb.ProxyDeleteSiteRequest) (*pb.ProxyDeleteSiteReply, error) {
	_, rep, err := s.deletesite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ProxyDeleteSiteReply), nil
}

func (s *grpcServer) CheckSitenameExists(ctx context.Context, req *pb.CheckSitenameExistsRequest) (*pb.CheckSitenameExistsReply, error) {
	_, rep, err := s.checksitenameexists.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CheckSitenameExistsReply), nil
}

func (s *grpcServer) GetSiteIDByUserID(ctx context.Context, req *pb.GetSiteIDByUserIDRequest) (*pb.GetSiteIDByUserIDReply, error) {
	_, rep, err := s.getsiteidbyuserid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetSiteIDByUserIDReply), nil
}

func (s *grpcServer) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigReply, error) {
	_, rep, err := s.getconfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetConfigReply), nil
}

func (s *grpcServer) UpdateConfig(ctx context.Context, req *pb.UpdateConfigRequest) (*pb.UpdateConfigReply, error) {
	_, rep, err := s.updateconfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateConfigReply), nil
}

func (s *grpcServer) GetThemes(ctx context.Context, req *pb.GetThemesRequest) (*pb.GetThemesReply, error) {
	_, rep, err := s.getthemes.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetThemesReply), nil
}

func (s *grpcServer) UpdateSiteTheme(ctx context.Context, req *pb.UpdateSiteThemeRequest) (*pb.UpdateSiteThemeReply, error) {
	_, rep, err := s.updatesitetheme.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateSiteThemeReply), nil
}

func (s *grpcServer) PostAbout(ctx context.Context, req *pb.PostAboutRequest) (*pb.PostAboutReply, error) {
	_, rep, err := s.postabout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PostAboutReply), nil
}

func (s *grpcServer) GetAbout(ctx context.Context, req *pb.GetAboutRequest) (*pb.GetAboutReply, error) {
	_, rep, err := s.getabout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAboutReply), nil
}

// NewSite
func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.CreateNewSiteRequest{Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewSiteRequest)
	return endpoint.NewSiteRequest{
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCNewSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewSiteResponse)
	return &pb.CreateNewSiteReply{
		SiteId: uint64(resp.SiteID),
		Err:    err2str(resp.Err),
	}, nil
}

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewSiteReply)
	return endpoint.NewSiteResponse{SiteID: uint(reply.SiteId), Err: str2err(reply.Err)}, nil
}

// DeleteSite
func encodeGRPCDeleteSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteSiteRequest)
	return &pb.ProxyDeleteSiteRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCDeleteSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ProxyDeleteSiteRequest)
	return endpoint.DeleteSiteRequest{
		SiteID: uint(req.SiteId),
	}, nil
}

func encodeGRPCDeleteSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeleteSiteResponse)
	return &pb.ProxyDeleteSiteReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCDeleteSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ProxyDeleteSiteReply)
	return endpoint.DeleteSiteResponse{Err: str2err(reply.Err)}, nil
}

// CheckSitenameExists
func encodeGRPCCheckSitenameExistsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CheckSitenameExistsRequest)
	return &pb.CheckSitenameExistsRequest{Sitename: req.Sitename}, nil
}

func decodeGRPCCheckSitenameExistsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CheckSitenameExistsRequest)
	return endpoint.CheckSitenameExistsRequest{
		Sitename: req.Sitename,
	}, nil
}

func encodeGRPCCheckSitenameExistsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CheckSitenameExistsResponse)
	return &pb.CheckSitenameExistsReply{
		Exists: resp.Exists,
		Err:    err2str(resp.Err),
	}, nil
}

func decodeGRPCCheckSitenameExistsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CheckSitenameExistsReply)
	return endpoint.CheckSitenameExistsResponse{Exists: reply.Exists, Err: str2err(reply.Err)}, nil
}

// GetSiteIDByUserID
func encodeGRPCGetSiteIDByUserIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.GetSiteIDByUserIDRequest{}, nil
}

func decodeGRPCGetSiteIDByUserIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoint.GetSiteIDByUserIDRequest{}, nil
}

func encodeGRPCGetSiteIDByUserIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetSiteIDByUserIDResponse)
	return &pb.GetSiteIDByUserIDReply{
		SiteId: uint64(resp.SiteID),
		Err:    err2str(resp.Err),
	}, nil
}

func decodeGRPCGetSiteIDByUserIDResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetSiteIDByUserIDReply)
	return endpoint.GetSiteIDByUserIDResponse{SiteID: uint(reply.SiteId), Err: str2err(reply.Err)}, nil
}

// GetConfig
func encodeGRPCGetConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetConfigRequest)
	return &pb.GetConfigRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCGetConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetConfigRequest)
	return endpoint.GetConfigRequest{
		SiteID: uint(req.SiteId),
	}, nil
}

func encodeGRPCGetConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetConfigResponse)
	return &pb.GetConfigReply{
		Config: resp.Config,
		Err:    err2str(resp.Err),
	}, nil
}

func decodeGRPCGetConfigResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetConfigReply)
	return endpoint.GetConfigResponse{Config: reply.Config, Err: str2err(reply.Err)}, nil
}

// UpdateConfig
func encodeGRPCUpdateConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateConfigRequest)
	return &pb.UpdateConfigRequest{SiteId: uint64(req.SiteID), Config: req.Config}, nil
}

func decodeGRPCUpdateConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateConfigRequest)
	return endpoint.UpdateConfigRequest{
		SiteID: uint(req.SiteId),
		Config: req.Config,
	}, nil
}

func encodeGRPCUpdateConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateConfigResponse)
	return &pb.UpdateConfigReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCUpdateConfigResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UpdateConfigReply)
	return endpoint.UpdateConfigResponse{Err: str2err(reply.Err)}, nil
}

// GetThemes
func encodeGRPCGetThemesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetThemesRequest)
	return &pb.GetThemesRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCGetThemesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetThemesRequest)
	return endpoint.GetThemesRequest{
		SiteID: uint(req.SiteId),
	}, nil
}

func encodeGRPCGetThemesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetThemesResponse)
	return &pb.GetThemesReply{
		Themes: resp.Themes,
		Err:    err2str(resp.Err),
	}, nil
}

func decodeGRPCGetThemesResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetThemesReply)
	return endpoint.GetThemesResponse{Themes: reply.Themes, Err: str2err(reply.Err)}, nil
}

// UpdateSiteTheme
func encodeGRPCUpdateSiteThemeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateSiteThemeRequest)
	return &pb.UpdateSiteThemeRequest{SiteId: uint64(req.SiteID), Theme: req.Theme}, nil
}

func decodeGRPCUpdateSiteThemeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateSiteThemeRequest)
	return endpoint.UpdateSiteThemeRequest{
		SiteID: uint(req.SiteId),
		Theme:  req.Theme,
	}, nil
}

func encodeGRPCUpdateSiteThemeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateSiteThemeResponse)
	return &pb.UpdateSiteThemeReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCUpdateSiteThemeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UpdateSiteThemeReply)
	return endpoint.UpdateSiteThemeResponse{Err: str2err(reply.Err)}, nil
}

// PostAbout
func encodeGRPCPostAboutRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.PostAboutRequest)
	return &pb.PostAboutRequest{SiteId: uint64(req.SiteID), Content: req.Content}, nil
}

func decodeGRPCPostAboutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PostAboutRequest)
	return endpoint.PostAboutRequest{
		SiteID:  uint(req.SiteId),
		Content: req.Content,
	}, nil
}

func encodeGRPCPostAboutResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.PostAboutResponse)
	return &pb.PostAboutReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCPostAboutResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.PostAboutReply)
	return endpoint.PostAboutResponse{Err: str2err(reply.Err)}, nil
}

// GetAbout
func encodeGRPCGetAboutRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetAboutRequest)
	return &pb.GetAboutRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCGetAboutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetAboutRequest)
	return endpoint.GetAboutRequest{
		SiteID: uint(req.SiteId),
	}, nil
}

func encodeGRPCGetAboutResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetAboutResponse)
	return &pb.GetAboutReply{
		Content: resp.Content,
		Err:     err2str(resp.Err),
	}, nil
}

func decodeGRPCGetAboutResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetAboutReply)
	return endpoint.GetAboutResponse{Content: reply.Content, Err: str2err(reply.Err)}, nil
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}
