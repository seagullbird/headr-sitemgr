package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/seagullbird/headr-sitemgr/endpoint"
	"net/http"
	"strconv"
)

type errorWrapper struct {
	Error string `json:"error"`
}

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	r := mux.NewRouter()

	// POST 	/sites/							add a site
	// DELETE	/sites/:id						remove the given site
	// POST     /is-sitename-exists 			check if sitename already exists

	r.Methods("POST").Path("/sites/").Handler(httptransport.NewServer(
		endpoints.NewSiteEndpoint,
		decodeHTTPNewSiteRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("DELETE").Path("/sites/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteSiteEndpoint,
		decodeHTTPDeleteSiteRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("POST").Path("/is-sitename-exists").Handler(httptransport.NewServer(
		endpoints.CheckSitenameExistsEndpoint,
		decodeHTTPCheckSitenameExistsRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return r
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

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPNewSiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.NewSiteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPDeleteSiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	return endpoint.DeleteSiteRequest{SiteId: uint(i)}, nil
}

func decodeHTTPCheckSitenameExistsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CheckSitenameExistsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
