package integration_test

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq/client"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-repoctl/transport"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
)

const (
	port = ":2345"
)

func TestIntegration(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var (
		servername = os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR")
		username   = "guest"
		passwd     = "guest"
	)
	dispatcher, err := dispatch.NewDispatcher(client.New(servername, username, passwd), logger)
	if err != nil {
		t.Fatal(err)
	}
	var (
		svc        = service.New(dispatcher, logger)
		endpoints  = endpoint.New(svc, logger)
		grpcServer = transport.NewGRPCServer(endpoints, logger)
	)

	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatal(err)
	}
	baseServer := grpc.NewServer()
	pb.RegisterRepoctlServer(baseServer, grpcServer)

	go baseServer.Serve(grpcListener)

	// Start GRPC client
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)
	siteID := uint(1)
	filename := "filename"
	content := "content"
	ctx := context.Background()
	// NewSite won't do anything but dispatch a message
	if err := client.NewSite(ctx, siteID, "theme"); err != nil {
		t.Fatal(err)
	}
	// WritePost will mkdir the post path
	if err := client.WritePost(ctx, siteID, filename, content); err != nil {
		t.Fatal(err)
	}
	// ReadPost
	readContent, err := client.ReadPost(ctx, siteID, filename)
	if err != nil {
		t.Fatal(err)
	}
	if readContent != content {
		t.Fatal(errors.New("ReadPost inconsistent with WritePost"))
	}
	// RemovePost
	if err := client.RemovePost(ctx, siteID, filename); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(service.PostPath(siteID, filename)); !os.IsNotExist(err) {
		t.Fatal(err)
	}
	// DeleteSite
	if err := client.DeleteSite(ctx, siteID); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(service.SitePath(siteID)); !os.IsNotExist(err) {
		t.Fatal(err)
	}
}
