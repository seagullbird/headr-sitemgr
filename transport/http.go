package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"net/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}
	m := http.NewServeMux()
	m.Handle("/create-new-site", httptransport.NewServer(
		endpoints.NewSiteEndpoint,
		decodeHTTPNewSiteRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/delete-site", httptransport.NewServer(
		endpoints.DeleteSiteEndpoint,
		decodeHTTPDeleteSiteRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return m
}

func err2code(err error) int {
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func decodeHTTPNewSiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.NewSiteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPDeleteSiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.DeleteSiteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
