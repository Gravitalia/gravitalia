package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Gravitalia/gravitalia/database"
	"github.com/Gravitalia/gravitalia/helpers"
	"github.com/Gravitalia/gravitalia/model"
)

var client = &http.Client{}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randomString generates a random character string with a predefined number
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// makeRequest allows to make requests and return the body
func makeRequest(url string, method string, reqBody io.Reader, authHeader string) ([]byte, error) {
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, errors.New("unable to make request")
	}

	if authHeader != "" {
		req.Header.Add("Authorization", authHeader)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("unable to make request")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("unable to read request")
	}

	return body, nil
}

// OAuth handles requests for connections, and will grant a Json Web Token
// or redirect the user to the public data sharing acceptance page.
func OAuth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.URL.Query().Has("state") && req.URL.Query().Has("code") {
		jsonEncoder := json.NewEncoder(w)

		val, err := database.Mem.Get(req.URL.Query().Get("state"))
		if err != nil || string(val.Value) != "ok" {
			w.WriteHeader(http.StatusBadRequest)
			jsonEncoder.Encode(model.RequestError{
				Error:   true,
				Message: "Invalid state",
			})
		} else {
			postBody, _ := json.Marshal(struct {
				ClientId     string `json:"client_id"`
				ClientSecret string `json:"client_secret"`
				Code         string `json:"code"`
				RedirectUri  string `json:"redirect_uri"`
			}{
				ClientId:     "suba",
				ClientSecret: os.Getenv("SECRET"),
				Code:         req.URL.Query().Get("code"),
				RedirectUri:  os.Getenv("REDIRECT_URL"),
			})

			body, err := makeRequest(os.Getenv("OAUTH_API")+"/oauth2/token", "POST", bytes.NewBuffer(postBody), "")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Internal error:" + err.Error(),
				})
				return
			}
			var data model.RequestError
			json.Unmarshal(body, &data)
			if data.Error {
				w.WriteHeader(http.StatusBadRequest)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Invalid code",
				})
				return
			}

			body, err = makeRequest(os.Getenv("OAUTH_API")+"/users/@me", "GET", nil, data.Message)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Internal error:" + err.Error(),
				})
				return
			}
			var user model.AuthaUser
			json.Unmarshal(body, &user)
			if user.Vanity == "" {
				w.WriteHeader(http.StatusBadRequest)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Invalid code",
				})
				return
			}

			token, err := helpers.CreateToken(user.Vanity)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Internal error:" + err.Error(),
				})
				return
			}

			// Check if account has been deleted 1 hour ago
			val, err = database.Mem.Get(user.Vanity + "-gd")
			if err != nil || string(val.Value) == "ok" {
				w.WriteHeader(http.StatusBadRequest)
				jsonEncoder.Encode(model.RequestError{
					Error:   true,
					Message: "Account deleted too soon",
				})
				return
			}

			database.CreateUser(user.Vanity)

			w.WriteHeader(http.StatusOK)
			jsonEncoder.Encode(model.RequestError{
				Error:   false,
				Message: token,
			})
		}
	} else {
		state := randomString(24)
		database.Set(state, "ok", 500)
		http.Redirect(w, req, os.Getenv("OAUTH_HOST")+"/oauth2/authorize?response_type=code&client_id=suba&scope=user&state="+state, http.StatusTemporaryRedirect)
	}
}
