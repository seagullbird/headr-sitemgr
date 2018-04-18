package endpoint_test

import (
	"bytes"
	"context"
	"github.com/go-errors/errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-repoctl/endpoint"
	svcmock "github.com/seagullbird/headr-repoctl/service/mock"
	"testing"
)

// When testing Set, just make sure its output and that of its internal service are consistent.
func TestSet(t *testing.T) {
	// Mocking Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)

	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	endpoints := endpoint.New(mockSvc, logger)

	dummyError := errors.New("dummy error")
	tests := []struct {
		name string
		rets map[string][]interface{}
	}{
		{"No Error", map[string][]interface{}{
			"NewSite":     {nil},
			"DeleteSite":  {nil},
			"WritePost":   {nil},
			"RemovePost":  {nil},
			"ReadPost":    {"string", nil},
			"WriteConfig": {nil},
			"ReadConfig":  {"string", nil},
			"UpdateAbout": {nil},
		}},
		{"Dummy Error", map[string][]interface{}{
			"NewSite":     {dummyError},
			"DeleteSite":  {dummyError},
			"WritePost":   {dummyError},
			"RemovePost":  {dummyError},
			"ReadPost":    {"", dummyError},
			"WriteConfig": {dummyError},
			"ReadConfig":  {"", dummyError},
			"UpdateAbout": {dummyError},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set EXPECTS
			mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["NewSite"]...).Times(2)
			mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(tt.rets["DeleteSite"]...).Times(2)
			mockSvc.EXPECT().WritePost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["WritePost"]...).Times(2)
			mockSvc.EXPECT().RemovePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["RemovePost"]...).Times(2)
			mockSvc.EXPECT().ReadPost(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["ReadPost"]...).Times(2)
			mockSvc.EXPECT().WriteConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["WriteConfig"]...).Times(2)
			mockSvc.EXPECT().ReadConfig(gomock.Any(), gomock.Any()).Return(tt.rets["ReadConfig"]...).Times(2)
			mockSvc.EXPECT().UpdateAbout(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["UpdateAbout"]...).Times(2)

			t.Run("NewSite", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				theme := "theme"
				setErr := endpoints.NewSite(ctx, siteID, theme)
				svcErr := mockSvc.NewSite(ctx, siteID, theme)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				setErr := endpoints.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			})
			t.Run("WritePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				content := "content"
				setErr := endpoints.WritePost(ctx, siteID, filename, content)
				svcErr := mockSvc.WritePost(ctx, siteID, filename, content)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			})
			t.Run("RemovePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				setErr := endpoints.RemovePost(ctx, siteID, filename)
				svcErr := mockSvc.RemovePost(ctx, siteID, filename)
				if setErr != svcErr {
					t.Fatal("setErr=", setErr, "svcErr=", svcErr)
				}
			})
			t.Run("ReadPost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				setOutput, setErr := endpoints.ReadPost(ctx, siteID, filename)
				svcOutput, svcErr := mockSvc.ReadPost(ctx, siteID, filename)
				if setOutput != svcOutput || setErr != svcErr {
					t.Fatal(setOutput, setErr)
				}
			})
			t.Run("WriteConfig", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				config := "config"
				setErr := endpoints.WriteConfig(ctx, siteID, config)
				svcErr := mockSvc.WriteConfig(ctx, siteID, config)
				if setErr != svcErr {
					t.Fatal("setErr=", setErr, "svcErr=", svcErr)
				}
			})
			t.Run("ReadConfig", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				setOutput, setErr := endpoints.ReadConfig(ctx, siteID)
				svcOutput, svcErr := mockSvc.ReadConfig(ctx, siteID)
				if setOutput != svcOutput || setErr != svcErr {
					t.Fatal(setOutput, setErr)
				}
			})
			t.Run("UpdateAbout", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				content := "content"
				setErr := endpoints.UpdateAbout(ctx, siteID, content)
				svcErr := mockSvc.UpdateAbout(ctx, siteID, content)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			})
		})
	}
}

// In fact this part is tested in grpc_test.TestGRPCTransport, dual here for good coverage report
func TestSetBadEndpoint(t *testing.T) {
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
	}

	expectedMsg := "dummy error"
	if err := endpoints.NewSite(context.Background(), 1, "theme"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.DeleteSite(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.WritePost(context.Background(), 1, "", ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.RemovePost(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.ReadPost(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.WriteConfig(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.ReadConfig(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.UpdateAbout(context.Background(), 1, ""); err.Error() != expectedMsg {
		t.Fatal(err)
	}
}
