package transport

import (
	"google.golang.org/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/service"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/pb"
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-errors/errors"
)

type grpcServer struct {
	newsite 	grpctransport.Handler
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

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint: newsiteEndpoint,
	}
}

func (s *grpcServer) NewSite(ctx context.Context, req *pb.CreateNewSiteRequest) (*pb.CreateNewSiteReply, error) {
	_, rep, err := s.newsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateNewSiteReply), nil
}

func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.CreateNewSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewSiteRequest)
	return endpoint.NewSiteRequest{
		Email: req.Email,
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCNewSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewSiteResponse)
	return &pb.CreateNewSiteReply{
		Err: resp.Err.Error(),
	}, nil
}

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewSiteReply)
	return endpoint.NewSiteResponse{Err: errors.New(reply.Err)}, nil
}