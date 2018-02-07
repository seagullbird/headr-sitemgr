package transport

import (
	"google.golang.org/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/pb"
	"context"
	"github.com/go-errors/errors"
	"github.com/seagullbird/headr-repoctl/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type grpcServer struct {
	newsite 	grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.RepoctlServer {
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
			"pb.Repoctl",
			"NewSite",
			encodeGRPCNewSiteRequest,
			decodeGRPCNewSiteResponse,
			pb.NewSiteReply{},
		).Endpoint()
	}

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint: newsiteEndpoint,
	}
}

func (s *grpcServer) NewSite(ctx context.Context, req *pb.NewSiteRequest) (*pb.NewSiteReply, error) {
	_, rep, err := s.newsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NewSiteReply), nil
}

func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.NewSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewSiteRequest)
	return endpoint.NewSiteRequest{
		Email: req.Email,
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCNewSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewSiteResponse)
	return &pb.NewSiteReply{
		Err: resp.Err.Error(),
	}, nil
}

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.NewSiteReply)
	return endpoint.NewSiteResponse{Err: errors.New(reply.Err)}, nil
}