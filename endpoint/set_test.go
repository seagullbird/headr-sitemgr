package endpoint_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	svcmock "github.com/seagullbird/headr-sitemgr/service/mock"
	"testing"
)

func TestSet(t *testing.T) {
	// Mocking Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	endpoints := endpoint.New(mockSvc, logger)

	// Login
	ctx := context.Background()
	accessToken := auth.Login()
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, accessToken)

	dummyError := errors.New("dummy error")
	tests := []struct {
		name string
		rets map[string][]interface{}
	}{
		{"No Error", map[string][]interface{}{
			"NewSite":             {uint(1), nil},
			"DeleteSite":          {nil},
			"CheckSitenameExists": {true, nil},
			"GetSiteIDByUserID":   {uint(1), nil},
			"GetConfig":           {"string", nil},
			"UpdateConfig":        {nil},
			"GetThemes":           {"themes", nil},
			"UpdateSiteTheme":     {nil},
			"PostAbout":           {nil},
			"GetAbout":            {"content", nil},
		}},
		{"Dummy Error", map[string][]interface{}{
			"NewSite":             {uint(0), dummyError},
			"DeleteSite":          {dummyError},
			"CheckSitenameExists": {false, dummyError},
			"GetSiteIDByUserID":   {uint(0), dummyError},
			"GetConfig":           {"", dummyError},
			"UpdateConfig":        {dummyError},
			"GetThemes":           {"", dummyError},
			"UpdateSiteTheme":     {dummyError},
			"PostAbout":           {dummyError},
			"GetAbout":            {"", dummyError},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set EXPECTS
			times := 2
			mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(tt.rets["NewSite"]...).Times(times)
			mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(tt.rets["DeleteSite"]...).Times(times)
			mockSvc.EXPECT().CheckSitenameExists(gomock.Any(), gomock.Any()).Return(tt.rets["CheckSitenameExists"]...).Times(times)
			mockSvc.EXPECT().GetSiteIDByUserID(gomock.Any()).Return(tt.rets["GetSiteIDByUserID"]...).Times(times)
			mockSvc.EXPECT().GetConfig(gomock.Any(), gomock.Any()).Return(tt.rets["GetConfig"]...).Times(times)
			mockSvc.EXPECT().UpdateConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["UpdateConfig"]...).Times(times)
			mockSvc.EXPECT().GetThemes(gomock.Any(), gomock.Any()).Return(tt.rets["GetThemes"]...).Times(times)
			mockSvc.EXPECT().UpdateSiteTheme(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["UpdateSiteTheme"]...).Times(times)
			mockSvc.EXPECT().PostAbout(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["PostAbout"]...).Times(times)
			mockSvc.EXPECT().GetAbout(gomock.Any(), gomock.Any()).Return(tt.rets["GetAbout"]...).Times(times)

			t.Run("NewSite", func(t *testing.T) {
				sitename := "sitename"
				setSiteID, setErr := endpoints.NewSite(ctx, sitename)
				svcSiteID, svcErr := mockSvc.NewSite(ctx, sitename)
				if setSiteID != svcSiteID || setErr != svcErr {
					t.Fatal("\nsetSiteID: ", setSiteID, "\nsetErr: ", setErr, "\nsvcSiteID: ", svcSiteID, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				siteID := uint(1)
				setErr := endpoints.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)

				}
			})
			t.Run("CheckSitenameExists", func(t *testing.T) {
				sitename := "sitename"
				setOutput, setErr := endpoints.CheckSitenameExists(ctx, sitename)
				svcOutput, svcErr := mockSvc.CheckSitenameExists(ctx, sitename)
				if setErr != svcErr {
					t.Fatal("\nsetOutput: ", setOutput, "\nsetErr: ", setErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetSiteIDByUserID", func(t *testing.T) {
				setSiteID, setErr := endpoints.GetSiteIDByUserID(ctx)
				svcSiteID, svcErr := mockSvc.GetSiteIDByUserID(ctx)
				if setErr != svcErr {
					t.Fatal("\nsetSiteID: ", setSiteID, "\nsetErr: ", setErr, "\nsvcSiteID: ", svcSiteID, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetConfig", func(t *testing.T) {
				siteID := uint(1)
				setOutput, setErr := endpoints.GetConfig(ctx, siteID)
				svcOutput, svcErr := mockSvc.GetConfig(ctx, siteID)
				if setErr != svcErr {
					t.Fatal("\nsetOutput: ", setOutput, "\nsetErr: ", setErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("UpdateConfig", func(t *testing.T) {
				siteID := uint(1)
				config := "config"
				setErr := endpoints.UpdateConfig(ctx, siteID, config)
				svcErr := mockSvc.UpdateConfig(ctx, siteID, config)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetThemes", func(t *testing.T) {
				siteID := uint(1)
				setThemes, setErr := endpoints.GetThemes(ctx, siteID)
				svcThemes, svcErr := mockSvc.GetThemes(ctx, siteID)
				if setErr != svcErr {
					t.Fatal("\nsetThemes: ", setThemes, "\nsetErr: ", setErr, "\nsvcThemes: ", svcThemes, "\nsvcErr: ", svcErr)

				}
			})
			t.Run("UpdateSiteTheme", func(t *testing.T) {
				siteID := uint(1)
				theme := "theme"
				setErr := endpoints.UpdateSiteTheme(ctx, siteID, theme)
				svcErr := mockSvc.UpdateSiteTheme(ctx, siteID, theme)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("PostAbout", func(t *testing.T) {
				siteID := uint(1)
				content := "content"
				setErr := endpoints.PostAbout(ctx, siteID, content)
				svcErr := mockSvc.PostAbout(ctx, siteID, content)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetAbout", func(t *testing.T) {
				siteID := uint(1)
				setOutput, setErr := endpoints.GetAbout(ctx, siteID)
				svcOutput, svcErr := mockSvc.GetAbout(ctx, siteID)
				if setErr != svcErr {
					t.Fatal("\nsetOutput: ", setOutput, "\nsetErr: ", setErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
		})
	}
}

// In fact this part is tested in grpc_test.TestGRPCTransport, dual here for good coverage report
func TestSetBadEndpoint(t *testing.T) {
	makeBadEndpoint := func(resp interface{}) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := resp.(endpoint.Failer)
			return nil, r.Failed()
		}
	}

	endpoints := endpoint.Set{
		NewSiteEndpoint:             makeBadEndpoint(endpoint.NewSiteResponse{SiteID: 1, Err: errors.New("dummy error")}),
		DeleteSiteEndpoint:          makeBadEndpoint(endpoint.DeleteSiteResponse{Err: errors.New("dummy error")}),
		CheckSitenameExistsEndpoint: makeBadEndpoint(endpoint.CheckSitenameExistsResponse{Exists: true, Err: errors.New("dummy error")}),
		GetSiteIDByUserIDEndpoint:   makeBadEndpoint(endpoint.GetSiteIDByUserIDResponse{SiteID: 1, Err: errors.New("dummy error")}),
		GetConfigEndpoint:           makeBadEndpoint(endpoint.GetConfigResponse{Config: "config", Err: errors.New("dummy error")}),
		UpdateConfigEndpoint:        makeBadEndpoint(endpoint.UpdateConfigResponse{Err: errors.New("dummy error")}),
		GetThemesEndpoint:           makeBadEndpoint(endpoint.GetThemesResponse{Err: errors.New("dummy error")}),
		UpdateSiteThemeEndpoint:     makeBadEndpoint(endpoint.UpdateSiteThemeResponse{Err: errors.New("dummy error")}),
		PostAboutEndpoint:           makeBadEndpoint(endpoint.PostAboutResponse{Err: errors.New("dummy error")}),
		GetAboutEndpoint:            makeBadEndpoint(endpoint.GetAboutResponse{Err: errors.New("dummy error")}),
	}

	expectedMsg := "dummy error"
	if _, err := endpoints.NewSite(context.Background(), "sitename"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.DeleteSite(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.CheckSitenameExists(context.Background(), "sitename"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetSiteIDByUserID(context.Background()); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetConfig(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.UpdateConfig(context.Background(), 1, "config"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetThemes(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.UpdateSiteTheme(context.Background(), 1, "theme"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.PostAbout(context.Background(), 1, "theme"); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetAbout(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
}
