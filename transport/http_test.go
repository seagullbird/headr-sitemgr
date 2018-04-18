package transport

import (
	"bytes"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	svcmock "github.com/seagullbird/headr-sitemgr/service/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPForbidden(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	server := httptest.NewServer(NewHTTPHandler(endpoints, logger))
	defer server.Close()
	client := &http.Client{}

	tests := []struct {
		name         string
		path         string
		method       string
		body         string
		expectedCode int
	}{
		{"DeleteSiteBadRouting", "/sites/a", "DELETE", "", http.StatusBadRequest},
		{"GetConfigBadRouting", "/sites/config/a", "GET", "", http.StatusBadRequest},
		{"UpdateConfigBadRouting", "/sites/config/a", "PUT", "{\"config\": \"config\"}", http.StatusBadRequest},
		{"GetThemesBadRouting", "/sites/themes/?aaa=1", "GET", "", http.StatusBadRequest},
		{"UpdateSiteThemeBadRouting", "/sites/a", "PATCH", "{\"theme\": \"theme\"}", http.StatusBadRequest},
		{"PostAboutBadRouting", "/sites/about/a", "PUT", "{\"content\": \"content\"}", http.StatusBadRequest},

		{"NewSite", "/sites/", "POST", "{\"sitename\": \"sitename\"}", http.StatusForbidden},
		{"DeleteSite", "/sites/1", "DELETE", "", http.StatusForbidden},
		{"CheckSitenameExists", "/is-sitename-exists", "POST", "{\"sitename\": \"sitename\"}", http.StatusForbidden},
		{"GetSiteIDByUserID", "/site-id/", "GET", "", http.StatusForbidden},
		{"GetConfig", "/sites/config/1", "GET", "", http.StatusForbidden},
		{"UpdateConfig", "/sites/config/1", "PUT", "{\"config\": \"config\"}", http.StatusForbidden},
		{"GetThemes", "/sites/themes/?site_id=1", "GET", "", http.StatusForbidden},
		{"UpdateSiteTheme", "/sites/1", "PATCH", "{\"theme\": \"theme\"}", http.StatusForbidden},
		{"PostAbout", "/sites/about/1", "PUT", "{\"content\": \"content\"}", http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(tt.body)
			req, err := http.NewRequest(tt.method, server.URL+tt.path, bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Error in creating %s to %s: %v", tt.method, tt.path, err)
			}
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error in %s to %s: %v", tt.method, tt.path, err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.expectedCode {
				t.Fatalf("Unexpected status code: %d\n Status code should be %d", resp.StatusCode, tt.expectedCode)
			}
		})
	}
}

func TestHTTP(t *testing.T) {
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
			"GetConfig":           {"config", nil},
			"UpdateConfig":        {nil},
			"GetThemes":           {"themes", nil},
			"UpdateSiteTheme":     {nil},
			"PostAbout":           {nil},
		},
		{
			"NewSite":             {uint(0), dummyError},
			"DeleteSite":          {dummyError},
			"CheckSitenameExists": {false, dummyError},
			"GetSiteIDByUserID":   {uint(0), dummyError},
			"GetConfig":           {"", dummyError},
			"UpdateConfig":        {dummyError},
			"GetThemes":           {"", dummyError},
			"UpdateSiteTheme":     {dummyError},
			"PostAbout":           {dummyError},
		},
	} {
		times := 1
		mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(rets["NewSite"]...).Times(times)
		mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(rets["DeleteSite"]...).Times(times)
		mockSvc.EXPECT().CheckSitenameExists(gomock.Any(), gomock.Any()).Return(rets["CheckSitenameExists"]...).Times(times)
		mockSvc.EXPECT().GetSiteIDByUserID(gomock.Any()).Return(rets["GetSiteIDByUserID"]...).Times(times)
		mockSvc.EXPECT().GetConfig(gomock.Any(), gomock.Any()).Return(rets["GetConfig"]...).Times(times)
		mockSvc.EXPECT().UpdateConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["UpdateConfig"]...).Times(times)
		mockSvc.EXPECT().GetThemes(gomock.Any(), gomock.Any()).Return(rets["GetThemes"]...).Times(times)
		mockSvc.EXPECT().UpdateSiteTheme(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["UpdateSiteTheme"]...).Times(times)
		mockSvc.EXPECT().PostAbout(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["PostAbout"]...).Times(times)
	}

	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	server := httptest.NewServer(NewHTTPHandler(endpoints, logger))
	defer server.Close()
	client := &http.Client{}

	// Login
	accessToken := auth.Login()
	// testcases
	tests := map[string][]struct {
		name             string
		path             string
		method           string
		body             string
		expectedCode     int
		expectedRespBody string
	}{
		"No Error": {
			{"NewSite", "/sites/", "POST", "{\"sitename\": \"sitename\"}", http.StatusOK, "{\"site_id\":1}"},
			{"DeleteSite", "/sites/1", "DELETE", "", http.StatusOK, "{}"},
			{"CheckSitenameExists", "/is-sitename-exists", "POST", "{\"sitename\": \"sitename\"}", http.StatusOK, "{\"exists\":true}"},
			{"GetSiteIDByUserID", "/site-id/", "GET", "", http.StatusOK, "{\"site_id\":1}"},
			{"GetConfig", "/sites/config/1", "GET", "", http.StatusOK, "{\"config\":\"config\"}"},
			{"UpdateConfig", "/sites/config/1", "PUT", "{\"config\": \"config\"}", http.StatusOK, "{}"},
			{"GetThemes", "/sites/themes/?site_id=1", "GET", "", http.StatusOK, "{\"themes\":\"themes\"}"},
			{"UpdateSiteTheme", "/sites/1", "PATCH", "{\"theme\": \"theme\"}", http.StatusOK, "{}"},
			{"PostAbout", "/sites/about/1", "PUT", "{\"content\": \"content\"}", http.StatusOK, "{}"},
		},
		"Dummy Error": {
			{"NewSite", "/sites/", "POST", "{\"sitename\": \"sitename\"}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"DeleteSite", "/sites/1", "DELETE", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"CheckSitenameExists", "/is-sitename-exists", "POST", "{\"sitename\": \"sitename\"}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"GetSiteIDByUserID", "/site-id/", "GET", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"GetConfig", "/sites/config/1", "GET", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"UpdateConfig", "/sites/config/1", "PUT", "{\"config\": \"config\"}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"GetThemes", "/sites/themes/?site_id=1", "GET", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"UpdateSiteTheme", "/sites/1", "PATCH", "{\"theme\": \"theme\"}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"PostAbout", "/sites/about/1", "PUT", "{\"content\": \"content\"}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
		},
	}

	// Start tests
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			for _, tt := range v {
				t.Run(tt.name, func(t *testing.T) {
					body := []byte(tt.body)
					req, err := http.NewRequest(tt.method, server.URL+tt.path, bytes.NewBuffer(body))
					if err != nil {
						t.Fatalf("Error in creating %s to %s: %v", tt.method, tt.path, err)
					}
					req.Header.Add("Content-Type", "application/json")
					req.Header.Add("Authorization", "Bearer "+accessToken)
					resp, err := client.Do(req)
					if err != nil {
						t.Fatalf("Error in %s to %s: %v", tt.method, tt.path, err)
					}
					defer resp.Body.Close()
					if resp.StatusCode != tt.expectedCode {
						t.Fatalf("Unexpected status code: %d\n Status code should be %d", resp.StatusCode, tt.expectedCode)
					}
					payload, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						t.Fatalf("Error in reading response body: %v", err)
					}
					respBody := strings.Trim(string(payload), "\n")
					if respBody != tt.expectedRespBody {
						t.Fatalf("Unexpected response body\nwant:\n%s\nget:\n%s\n", tt.expectedRespBody, respBody)
					}
				})
			}
		})
	}
}
