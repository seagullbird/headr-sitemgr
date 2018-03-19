package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-sitemgr/config"
	"net/http"
	"time"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func AuthMiddleware() endpoint.Middleware {
	getPemCert := func(token *stdjwt.Token) (string, error) {
		type JSONWebKeys struct {
			Kty string   `json:"kty"`
			Kid string   `json:"kid"`
			Use string   `json:"use"`
			N   string   `json:"n"`
			E   string   `json:"e"`
			X5c []string `json:"x5c"`
		}

		type Jwks struct {
			Keys []JSONWebKeys `json:"keys"`
		}

		cert := ""
		resp, err := http.Get(config.AUTH0_DOMAIN + "/.well-known/jwks.json")

		if err != nil {
			return cert, err
		}
		defer resp.Body.Close()

		var jwks = Jwks{}
		err = json.NewDecoder(resp.Body).Decode(&jwks)

		if err != nil {
			return cert, err
		}

		x5c := jwks.Keys[0].X5c
		for k, v := range x5c {
			if token.Header["kid"] == jwks.Keys[k].Kid {
				cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
			}
		}

		if cert == "" {
			err := errors.New("Unable to find appropriate key.")
			return cert, err
		}

		return cert, nil
	}

	keyFunc := func(token *stdjwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := config.AUTH0_AUDIENCE
		checkAud := token.Claims.(stdjwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("Invalid audience.")
		}
		// Verify 'iss' claim
		iss := config.AUTH0_DOMAIN + "/"
		checkIss := token.Claims.(stdjwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := stdjwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}

	return jwt.NewParser(keyFunc, stdjwt.SigningMethodRS256, jwt.MapClaimsFactory)
}

func Middlewares(e endpoint.Endpoint, logger log.Logger) endpoint.Endpoint {
	chain := endpoint.Chain(AuthMiddleware(), LoggingMiddleware(logger))
	return chain(e)
}
