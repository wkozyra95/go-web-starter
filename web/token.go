package web

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	conf "github.com/wkozyra95/go-web-starter/config"
	"gopkg.in/mgo.v2/bson"
)

type contextKeyType string

const contextUserId contextKeyType = "userId"

type jwtProvider struct {
	jwtKey []byte
	header string
}

func newJwtProvider(config conf.Config) (jwtProvider, error) {
	const keySize = 64
	jwtKey := make([]byte, keySize)
	_, err := rand.Read(jwtKey)
	if err != nil {
		return jwtProvider{}, err
	}
	return jwtProvider{
		jwtKey: jwtKey,
		header: "X-Auth-Token",
	}, nil
}

func (jp jwtProvider) generate(id bson.ObjectId) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(jp.jwtKey)
}

func (jp jwtProvider) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get(jp.header)
		if token == "" {
			log.Warn("Missing auth token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parsed, parseErr := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
			return jp.jwtKey, nil
		})
		if parseErr != nil {
			log.Warn("Unable to parse token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, assertTypeOk := parsed.Claims.(jwt.MapClaims)
		if parsed.Valid && assertTypeOk {
			idKey := contextUserId
			idVal := claims["id"]

			ctx := context.WithValue(r.Context(), idKey, idVal)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Warn("Token is not valid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
