// Package account manage user info and valid/sign process for the Info
// IMPORTANT: It uses HS256 for signing JWT. For security use RSA insteed
package account

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

//InfoContextKey for get information from the requet context
var InfoContextKey = "AccountInfo"

// Repo for the accounts
type Repo struct {
	SigningKey []byte
}

// Info for the account
type Info struct {
	UID   string
	Chips int
	Bet   int
	jwt.StandardClaims
}

// InfoFromContext fetch Info from the request context
func (r *Repo) InfoFromContext(ctx context.Context) (*Info, error) {
	v, ok := ctx.Value(&InfoContextKey).(*Info)
	if !ok {
		return nil, errors.New("Account info not found")
	}
	return v, nil
}

// SignInfo  with signing key. Returns jwt token string
func (r *Repo) SignInfo(i *Info) (string, error) {
	// Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), i)
	// token -> string. Only server knows this secret
	tokenstring, err := token.SignedString(r.SigningKey)
	if err != nil {
		return "", errors.Wrap(err, "Couldn't sign JWT")
	}
	return tokenstring, nil
}

// Middleware for processing user account Info
// IT decode jwt and store metadata into context
// Return 403 for not valid tokens
func (r *Repo) Middleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		buf, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		defer req.Body.Close() // nolint: errcheck
		reqB := map[string]string{}
		err = json.Unmarshal(buf, &reqB)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString, ok := reqB["jwt"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var info Info
		token, err := jwt.ParseWithClaims(tokenString, &info,
			func(token *jwt.Token) (interface{}, error) {
				return r.SigningKey, nil
			})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), &InfoContextKey, &info)
		h.ServeHTTP(w, req.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
