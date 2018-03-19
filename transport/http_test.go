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
	logger := log.NewNopLogger()
	emptyEndpoints := endpoint.New(service.EmptyService{}, logger)
	server := httptest.NewServer(NewHTTPHandler(emptyEndpoints, logger))
	defer server.Close()
	client := &http.Client{}
	right_access_token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlFqZzVSVEZHUVRCQ09VSkNOalpGTjBZM016RXhNRUkxTUVORU5rTkNRVVkwTXpGQlFrTXlOZyJ9.eyJodHRwOi8vaGVhZHIvZW1haWwiOiJoZWxsb0BxcS5jb20iLCJpc3MiOiJodHRwczovL2hlYWRyLmF1dGgwLmNvbS8iLCJzdWIiOiJhdXRoMHw1YWFjOTljZWVjZDkwZTAyNjMyM2UyYjAiLCJhdWQiOlsiaHR0cHM6Ly9hcGkuaGVhZHIuaW8iLCJodHRwczovL2hlYWRyLmF1dGgwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE1MjE0Mzk4ODMsImV4cCI6MTUyMTQ0NzA4MywiYXpwIjoiVlVwWEdiZzdLaHZjS0tjQ09QdVBHT3BITXkyc1l6UmQiLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIn0.W12QN1364mUVW5s2Xpq_kp3mEJ64w_EQXozOhPG4sZCUNN1FdWXGgRqrR0GwQ9S6H20zcLoxiTIOaXwLhpNN3-QocHKu7X4-2NaFZJ8o-_vT7ryi6cBGr3tHloY9LVE5wl4FqRO1BJBmZn0jGGrT2039-puMdZ2n6xda26-6Bw5ai-KQikAC5Wx_I3XE3Q5jD1-KMlrJzKBuuxcm8z94DFt_uAqv8TBUc83S5hhvrYavck8D_LO3G74dpzA5ELTFDLrSlG6dNqX1UtuSu8aAHGi0J-e0Rt3526VadujhcafDUPDXFZsVkSycQi23tQPwP4i4NRhEOclsBzTszZKrsQ"
	body := []byte("{\"sitename\": \"test-site\"}")

	req, err := http.NewRequest("POST", server.URL+"/is-sitename-exists", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST to /is-sitename-exists: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+right_access_token)
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to /is-sitename-exists: %v", err)
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error in reading response body: %v", err)
	}

	fmt.Println(string(payload))
}
