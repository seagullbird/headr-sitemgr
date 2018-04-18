package transport_test

import (
	"github.com/golang/mock/gomock"
	svcmock "github.com/seagullbird/headr-repoctl/service/mock"
	"testing"

	"context"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/transport"
	"google.golang.org/grpc"
	"net"
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
	pb.RegisterRepoctlServer(baseServer, grpcServer)
	baseServer.Serve(grpcListener)
}

// This Test tests only application level error handling;
// Application error means an error that is returned by the service itself;
// For example, if the client requested information of an unknown user, the service will respond with a
// "User Not Found" error, which is usually customized by the service programmer.
// So this Test mocks the service, which returns two types of response: No Error and Dummy Error;
// I want to make sure that what the client receives is exactly the same as what the service returns. (nil or dummy error)
func TestGRPCApplication(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	// Set mock service expectations
	dummyError := errors.New("dummy error")
	for _, rets := range []map[string][]interface{}{
		{
			"NewSite":     {nil},
			"DeleteSite":  {nil},
			"WritePost":   {nil},
			"RemovePost":  {nil},
			"ReadPost":    {"string", nil},
			"WriteConfig": {nil},
			"ReadConfig":  {"string", nil},
			"UpdateAbout": {nil},
			"ReadAbout":   {"string", nil},
		},
		{
			"NewSite":     {dummyError},
			"DeleteSite":  {dummyError},
			"WritePost":   {dummyError},
			"RemovePost":  {dummyError},
			"ReadPost":    {"", dummyError},
			"WriteConfig": {dummyError},
			"ReadConfig":  {"", dummyError},
			"UpdateAbout": {dummyError},
			"ReadAbout":   {"", dummyError},
		},
	} {
		times := 2
		mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["NewSite"]...).Times(times)
		mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(rets["DeleteSite"]...).Times(times)
		mockSvc.EXPECT().WritePost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["WritePost"]...).Times(times)
		mockSvc.EXPECT().RemovePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["RemovePost"]...).Times(times)
		mockSvc.EXPECT().ReadPost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["ReadPost"]...).Times(times)
		mockSvc.EXPECT().WriteConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["WriteConfig"]...).Times(times)
		mockSvc.EXPECT().ReadConfig(gomock.Any(), gomock.Any()).Return(rets["ReadConfig"]...).Times(times)
		mockSvc.EXPECT().UpdateAbout(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["UpdateAbout"]...).Times(times)
		mockSvc.EXPECT().ReadAbout(gomock.Any(), gomock.Any()).Return(rets["ReadAbout"]...).Times(times)
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
				siteID := uint(1)
				ctx := context.Background()
				theme := "theme"
				clientErr := client.NewSite(ctx, siteID, theme)
				svcErr := mockSvc.NewSite(ctx, siteID, theme)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				clientErr := client.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("WritePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				content := "content"
				clientErr := client.WritePost(ctx, siteID, filename, content)
				svcErr := mockSvc.WritePost(ctx, siteID, filename, content)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("RemovePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				clientErr := client.RemovePost(ctx, siteID, filename)
				svcErr := mockSvc.RemovePost(ctx, siteID, filename)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("ReadPost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				clientOutput, clientErr := client.ReadPost(ctx, siteID, filename)
				svcOutput, svcErr := mockSvc.ReadPost(ctx, siteID, filename)
				if clientOutput != svcOutput || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("WriteConfig", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				config := "filename"
				clientErr := client.WriteConfig(ctx, siteID, config)
				svcErr := mockSvc.WriteConfig(ctx, siteID, config)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("ReadConfig", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				clientOutput, clientErr := client.ReadConfig(ctx, siteID)
				svcOutput, svcErr := mockSvc.ReadConfig(ctx, siteID)
				if clientOutput != svcOutput || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("UpdateAbout", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				content := "content"
				clientErr := client.UpdateAbout(ctx, siteID, content)
				svcErr := mockSvc.UpdateAbout(ctx, siteID, content)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("ReadAbout", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				clientOutput, clientErr := client.ReadAbout(ctx, siteID)
				svcOutput, svcErr := mockSvc.ReadAbout(ctx, siteID)
				if clientOutput != svcOutput || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
		})
	}

	baseServer.Stop()
}

// This Test on the other hand focuses on transport level error,
// such as "rpc error: code = Unknown desc = transport is closing", which are defined by grpc but not the service author.
// I want to make sure that when a transport level error occurs, the client will receive the right transport error.
func TestGRPCTransport(t *testing.T) {
	makeBadEndpoint := func() kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return nil, errors.New("dummy error")
		}
	}

	endpoints := endpoint.Set{
		NewSiteEndpoint:     makeBadEndpoint(),
		DeleteSiteEndpoint:  makeBadEndpoint(),
		WritePostEndpoint:   makeBadEndpoint(),
		RemovePostEndpoint:  makeBadEndpoint(),
		ReadPostEndpoint:    makeBadEndpoint(),
		WriteConfigEndpoint: makeBadEndpoint(),
		ReadConfigEndpoint:  makeBadEndpoint(),
		UpdateAboutEndpoint: makeBadEndpoint(),
		ReadAboutEndpoint:   makeBadEndpoint(),
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
	if err := client.NewSite(context.Background(), 1, "theme"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.DeleteSite(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.WritePost(context.Background(), 1, "", ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.RemovePost(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.ReadPost(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.WriteConfig(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.ReadConfig(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.UpdateAbout(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.ReadAbout(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}

	baseServer.Stop()
}
