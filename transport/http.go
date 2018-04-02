package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
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

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(jwt.HTTPToContext()),
	}

	r := mux.NewRouter()

	// POST 	/sites/							add a site
	// DELETE	/sites/:id						remove the given site
	// POST     /is-sitename-exists 			check if sitename already exists
	// GET		/site-id						get a user's site's id for the given user id
	// GET		/sites/config/:id				get a site's config by site id
	// PUT		/sites/config/:id				update a site's config

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
	r.Methods("GET").Path("/site-id/").Handler(httptransport.NewServer(
		endpoints.GetSiteIDByUserIDEndpoint,
		decodeHTTPGetSiteIDByUserIDRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("GET").Path("/sites/config/{id}").Handler(httptransport.NewServer(
		endpoints.GetConfigEndpoint,
		decodeHTTPGetConfigRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("PUT").Path("/sites/config/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateConfigEndpoint,
		decodeHTTPUpdateConfigRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return r
}

func err2code(err error) int {
	switch err {
	case jwt.ErrTokenContextMissing, jwt.ErrTokenExpired, jwt.ErrTokenInvalid, jwt.ErrTokenMalformed, jwt.ErrTokenNotActive, jwt.ErrUnexpectedSigningMethod:
		return http.StatusForbidden
	case ErrBadRouting:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
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
	return endpoint.DeleteSiteRequest{SiteID: uint(i)}, nil
}

func decodeHTTPGetSiteIDByUserIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// UserID will not be set here but extracted from access_token after access_token is verified in endpoint.AuthMiddleware
	return endpoint.GetSiteIDByUserIDRequest{}, nil
}

func decodeHTTPGetConfigRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	return endpoint.GetConfigRequest{SiteID: uint(i)}, nil
}

func decodeHTTPUpdateConfigRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	var payload struct {
		Config string `json:"config"`
	}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, ErrBadRouting
	}
	return endpoint.UpdateConfigRequest{SiteID: uint(i), Config: payload.Config}, nil
}

func decodeHTTPCheckSitenameExistsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CheckSitenameExistsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
