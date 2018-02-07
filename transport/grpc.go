package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/pb"
	"context"
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

func (s *grpcServer) NewSite(ctx context.Context, req *pb.NewSiteRequest) (*pb.NewSiteReply, error) {
	_, rep, err := s.newsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NewSiteReply), nil
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