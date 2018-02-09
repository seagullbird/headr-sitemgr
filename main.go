package main

import (
	"os"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"fmt"
	repoctltransport "github.com/seagullbird/headr-repoctl/transport"
	"net"
	"github.com/seagullbird/headr-sitemgr/config"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/service"
	"github.com/seagullbird/headr-sitemgr/transport"
	"github.com/seagullbird/headr-sitemgr/pb"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Repoctl gRPC service
	conn, err := grpc.Dial("repoctl:2018", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	repoctlsvc := repoctltransport.NewGRPCClient(conn, logger)

	var (
		service = service.New(repoctlsvc, logger)
		endpoints = endpoint.New(service, logger)
		grpcServer = transport.NewGRPCServer(endpoints, logger)
	)

	grpcListener, err := net.Listen("tcp", config.PORT)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "gRPC", "addr", config.PORT)
	baseServer := grpc.NewServer()
	pb.RegisterSitemgrServer(baseServer, grpcServer)

	baseServer.Serve(grpcListener)
}