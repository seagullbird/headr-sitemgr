package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq/client"
	"github.com/seagullbird/headr-common/mq/dispatch"
	repoctltransport "github.com/seagullbird/headr-repoctl/transport"
	"github.com/seagullbird/headr-sitemgr/config"
	"github.com/seagullbird/headr-sitemgr/db"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/pb"
	"github.com/seagullbird/headr-sitemgr/service"
	"github.com/seagullbird/headr-sitemgr/transport"
	"google.golang.org/grpc"
	"net"
	"os"
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

	// mq dispatcher
	var (
		servername = os.Getenv("RABBITMQ_SERVER")
		username   = os.Getenv("RABBITMQ_USER")
		passwd     = os.Getenv("RABBITMQ_PASS")
	)
	dispatcher, err := dispatch.NewDispatcher(client.New(servername, username, passwd), logger)
	if err != nil {
		logger.Log("error_desc", "dispatch.NewDispatcher failed", "error", err)
		return
	}

	// database
	store := db.New(logger)
	// repoctl service
	repoctlsvc := repoctltransport.NewGRPCClient(conn, logger)
	var (
		service    = service.New(repoctlsvc, logger, dispatcher, store)
		endpoints  = endpoint.New(service, logger)
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
