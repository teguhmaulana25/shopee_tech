package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type res struct {
	Code    int         `json:"code"`    // code error for more explanation ex: 20003
	Message string      `json:"message"` // must as verbose as possible ex: failed Authenticate
	Data    interface{} `json:"data"`
}

// JwtAuth is verify token from client
func JwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				m := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
				log.Debug(m)
				return nil, errors.New(m)
			}

			return []byte(os.Getenv("JWTKEY")), nil
		})

		if token != nil && err == nil {
			log.Debug("JWT middleware passed")
			next.ServeHTTP(w, r)
		} else {
			log.Debug(err.Error())
			result := struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{
				"You're not authorized",
				err.Error(),
			}
			jsonData, err := json.Marshal(result)
			if err != nil {
				log.Error("Failed marshal data")
				jsonData = nil
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonData)
		}
	}
}
