package transport

import (
	"bytes"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"github.com/seagullbird/headr-sitemgr/service"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPHandler(t *testing.T) {
	//var logger log.Logger
	//{
	//	logger = log.NewLogfmtLogger(os.Stderr)
	//	logger = log.With(logger, "caller", log.DefaultCaller)
	//}
	logger := log.NewNopLogger()
	emptyEndpoints := endpoint.New(service.EmptyService{}, logger)

	server := httptest.NewServer(NewHTTPHandler(emptyEndpoints, logger))
	defer server.Close()
	client := &http.Client{}
	body := []byte("{\"sitename\": \"test-site\"}")

	// Test invalid access token
	invalidAccessToken := "invalid.token"
	req, err := http.NewRequest("POST", server.URL+"/is-sitename-exists", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST to /is-sitename-exists: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+invalidAccessToken)
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to /is-sitename-exists: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Unexpected status code: %d\n Status code should be %d", resp.StatusCode, http.StatusForbidden)
	}

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error in reading response body: %v", err)
	}

	fmt.Println(string(payload))
}
