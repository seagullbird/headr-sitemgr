package transport

import (
	"context"
	"github.com/go-errors/errors"
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
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.SitemgrServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
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
	}
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	var newsiteEndpoint kitendpoint.Endpoint
	{
		newsiteEndpoint = grpctransport.NewClient(
			conn,
			"pb.Sitemgr",
			"NewSite",
			encodeGRPCNewSiteRequest,
			decodeGRPCNewSiteResponse,
			pb.CreateNewSiteReply{},
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
		).Endpoint()
	}
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint:             newsiteEndpoint,
		DeleteSiteEndpoint:          deletesiteEndpoint,
		CheckSitenameExistsEndpoint: checksitenameexistsEndpoint,
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

// NewSite
func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.CreateNewSiteRequest{UserId: uint64(req.UserId), Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewSiteRequest)
	return endpoint.NewSiteRequest{
		UserId:   uint(req.UserId),
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCNewSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewSiteResponse)
	return &pb.CreateNewSiteReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewSiteReply)
	return endpoint.NewSiteResponse{Err: str2err(reply.Err)}, nil
}

// DeleteSite
func encodeGRPCDeleteSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteSiteRequest)
	return &pb.ProxyDeleteSiteRequest{SiteId: uint64(req.SiteId)}, nil
}

func decodeGRPCDeleteSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ProxyDeleteSiteRequest)
	return endpoint.DeleteSiteRequest{
		SiteId: uint(req.SiteId),
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
