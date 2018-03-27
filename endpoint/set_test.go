package endpoint_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
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
		}},
		{"Dummy Error", map[string][]interface{}{
			"NewSite":             {uint(0), dummyError},
			"DeleteSite":          {dummyError},
			"CheckSitenameExists": {false, dummyError},
			"GetSiteIDByUserID":   {uint(0), dummyError},
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
		})
	}
}
