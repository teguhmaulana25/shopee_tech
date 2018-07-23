package controller

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var errUsernameInvalid = errors.New("Username is invalid")
var errPasswordInvalid = errors.New("Password is invalid")
var errGenerateToken = errors.New("Generate token failed")
var errHasingPasswordFailed = errors.New("Hashing password failed")

type users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtToken struct {
	Token string `json:"token"`
}

// UserController represents the controller for operating on the User resource
type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (UserController) Auth(w http.ResponseWriter, r *http.Request) {
	users := users{}
	role := "admin"

	// Read the body of the request.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panic(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panic(err)
	}
	// Convert the JSON in the request to a Go type.
	if err := json.Unmarshal(body, &users); err != nil {
		log.Info("unmarshal")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panic(err)
		}
		response := res{
			Code:    500,
			Message: err.Error(),
		}
		renderJSON(w, response, http.StatusUnprocessableEntity)
		return
	}

	err = checkUsername(w, users.Username)
	if err != nil {
		response := res{
			Code:    500,
			Message: err.Error(),
		}
		renderJSON(w, response, http.StatusOK)
		return
	}
	err = checkPassword(w, []byte(users.Password))
	if err != nil {
		response := res{
			Code:    500,
			Message: err.Error(),
		}
		renderJSON(w, response, http.StatusOK)
		return
	}

	data, err := generateJWT(role)
	message := "Success"
	if err != nil {
		message = err.Error()
	}
	response := res{
		Code:    200,
		Message: message,
		Data:    data,
	}
	renderJSON(w, response, http.StatusOK)
}

func checkUsername(w http.ResponseWriter, username string) error {
	dummyUsername := "teguh"
	if username != dummyUsername {
		return errUsernameInvalid
	}

	return nil

}

func checkPassword(w http.ResponseWriter, password []byte) error {
	dummyPassword := "admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(dummyPassword), 14)
	if err != nil {
		log.Error("Generate bcrypt failed")
		return errHasingPasswordFailed
	}
	comparePassword := bcrypt.CompareHashAndPassword(hash, password)
	if comparePassword != nil {
		return errPasswordInvalid
	}

	return nil
}

func generateJWT(role string) (jwtToken, error) {
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "shopee.co.id",
		"exp":  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
		"role": role,
	})
	token, err := sign.SignedString([]byte(os.Getenv("JWTKEY")))
	if err != nil {
		return jwtToken{}, errGenerateToken
	}
	return jwtToken{
			token,
		},
		nil
}
