package service

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"

	"github.com/les-cours/gateway/env"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func decodeToken(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("authorization")
		tokenSplit := strings.Split(bearerToken, " ")
		if len(tokenSplit) != 2 {
			h.ServeHTTP(w, r)
			return
		}

		req, err := http.NewRequest("POST", env.Settings.AuthAPIEndPoint, nil)
		if err != nil {
			log.Println(err)
			h.ServeHTTP(w, r)
			return
		}

		req.Header.Set("Authorization", bearerToken)
		req.Header.Set("Origin", "https://"+env.Settings.HttpHost+":"+env.Settings.HttpPort)
		client := http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Println(err)
			h.ServeHTTP(w, r)
			return
		}

		message := ErrorMessage{}
		err = json.NewDecoder(response.Body).Decode(&message)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		if message.Message != "Token Valid" {
			h.ServeHTTP(w, r)
			return
		}

		user := UserToken{}
		_, _, err = new(jwt.Parser).ParseUnverified(tokenSplit[1], &user)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		h.ServeHTTP(w, r.WithContext(
			context.WithValue(r.Context(), "user", user),
		))
	})
}
