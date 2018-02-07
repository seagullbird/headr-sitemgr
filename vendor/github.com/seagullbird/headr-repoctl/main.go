package main

import (
	"os"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-repoctl/mq_helper"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/transport"
	"net"
	"google.golang.org/grpc"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/config"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var dispatcher = mq_helper.NewDispatcher("newsite")

	var (
		service = service.New(dispatcher, logger)
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
	pb.RegisterRepoctlServer(baseServer, grpcServer)

	baseServer.Serve(grpcListener)
}
