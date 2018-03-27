package transport_test

import (
	"context"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/pb"
	svcmock "github.com/seagullbird/headr-sitemgr/service/mock"
	"github.com/seagullbird/headr-sitemgr/transport"
	"google.golang.org/grpc"
	"net"
	"testing"
)

const (
	port = ":1234"
)

func startServer(t *testing.T, baseServer *grpc.Server, endpoints endpoint.Set, logger log.Logger) {
	grpcServer := transport.NewGRPCServer(endpoints, logger)
	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatal(err)
	}
	pb.RegisterSitemgrServer(baseServer, grpcServer)
	baseServer.Serve(grpcListener)
}

func TestGRPCApplication(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	// Set mock service expectations
	dummyError := errors.New("dummy error")
	for _, rets := range []map[string][]interface{}{
		{
			"NewSite":             {uint(1), nil},
			"DeleteSite":          {nil},
			"CheckSitenameExists": {true, nil},
			"GetSiteIDByUserID":   {uint(1), nil},
		},
		{
			"NewSite":             {uint(0), dummyError},
			"DeleteSite":          {dummyError},
			"CheckSitenameExists": {false, dummyError},
			"GetSiteIDByUserID":   {uint(0), dummyError},
		},
	} {
		times := 2
		mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(rets["NewSite"]...).Times(times)
		mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(rets["DeleteSite"]...).Times(times)
		mockSvc.EXPECT().CheckSitenameExists(gomock.Any(), gomock.Any()).Return(rets["CheckSitenameExists"]...).Times(times)
		mockSvc.EXPECT().GetSiteIDByUserID(gomock.Any()).Return(rets["GetSiteIDByUserID"]...).Times(times)
	}

	// Start GRPC server with the mock service
	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	baseServer := grpc.NewServer()
	go startServer(t, baseServer, endpoints, logger)

	// Start GRPC client
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)
	// Login
	ctx := context.Background()
	accessToken := auth.Login()
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, accessToken)

	// testcases
	tests := []struct {
		name   string
		judger func(err1, err2 error) bool
	}{
		{
			"No Error",
			func(err1, err2 error) bool {
				if err1 != nil || err2 != nil {
					return false
				}
				return true
			},
		},
		{
			"Dummy Error",
			func(err1, err2 error) bool {
				if err1.Error() != "dummy error" || err2.Error() != "dummy error" {
					return false
				}
				return true
			},
		},
	}

	// Start tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("NewSite", func(t *testing.T) {
				sitename := "sitename"
				clientSiteID, clientErr := client.NewSite(ctx, sitename)
				svcSiteID, svcErr := mockSvc.NewSite(ctx, sitename)
				if clientSiteID != svcSiteID || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientSiteID: ", clientSiteID, "\nclientErr: ", clientErr, "\nsvcSiteID: ", svcSiteID, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				siteID := uint(1)
				clientErr := client.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)

				}
			})
			t.Run("CheckSitenameExists", func(t *testing.T) {
				sitename := "sitename"
				clientOutput, clientErr := client.CheckSitenameExists(ctx, sitename)
				svcOutput, svcErr := mockSvc.CheckSitenameExists(ctx, sitename)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetSiteIDByUserID", func(t *testing.T) {
				clientSiteID, clientErr := client.GetSiteIDByUserID(ctx)
				svcSiteID, svcErr := mockSvc.GetSiteIDByUserID(ctx)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientSiteID: ", clientSiteID, "\nclientErr: ", clientErr, "\nsvcSiteID: ", svcSiteID, "\nsvcErr: ", svcErr)
				}
			})
		})
	}

	baseServer.Stop()
}

func TestGRPCTransport(t *testing.T) {
	makeBadEndpoint := func() kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return nil, errors.New("dummy error")
		}
	}

	endpoints := endpoint.Set{
		NewSiteEndpoint:             makeBadEndpoint(),
		DeleteSiteEndpoint:          makeBadEndpoint(),
		CheckSitenameExistsEndpoint: makeBadEndpoint(),
		GetSiteIDByUserIDEndpoint:   makeBadEndpoint(),
	}
	baseServer := grpc.NewServer()
	go startServer(t, baseServer, endpoints, log.NewNopLogger())

	// Start GRPC client
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)
	expectedMsg := "rpc error: code = Unknown desc = dummy error"
	if _, err := client.NewSite(context.Background(), "sitename"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.DeleteSite(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.CheckSitenameExists(context.Background(), "sitename"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.GetSiteIDByUserID(context.Background()); err.Error() != expectedMsg {
		t.Fatal(err)
	}

	baseServer.Stop()
}
