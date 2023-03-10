package router

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Gravitalia/gravitalia/database"
	"github.com/Gravitalia/gravitalia/helpers"
	"github.com/Gravitalia/gravitalia/model"
)

func Users(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)

	username := strings.TrimPrefix(req.URL.Path, "/users/")
	if username == "@me" && req.Header.Get("authorization") != "" {
		vanity, err := helpers.CheckToken(req.Header.Get("authorization"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			jsonEncoder.Encode(model.RequestError{
				Error:   true,
				Message: "Invalid token",
			})
			return
		}
		username = vanity
	}

	stats, err := database.GetUserStats(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(model.RequestError{
			Error:   true,
			Message: "Invalid user",
		})
		return
	}

	posts, err := database.GetUserPost(username, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(model.RequestError{
			Error:   true,
			Message: "Invalid user",
		})
		return
	}

	jsonEncoder.Encode(struct {
		Followers int64        `json:"followers"`
		Following int64        `json:"following"`
		Posts     []model.Post `json:"posts"`
	}{
		Followers: stats.Followers,
		Following: stats.Following,
		Posts:     posts,
	})
}

// Delete allows users to delete their account
func Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)

	vanity := ""
	var err error

	if req.Header.Get("authorization") == "" {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(model.RequestError{
			Error:   true,
			Message: "Invalid token",
		})
		return
	} else if req.Header.Get("authorization") == os.Getenv("GLOBAL_AUTH") {
		vanity = req.URL.Query().Get("user")
	} else {
		vanity, err = helpers.CheckToken(req.Header.Get("authorization"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			jsonEncoder.Encode(model.RequestError{
				Error:   true,
				Message: "Invalid token",
			})
			return
		}
	}

	_, err = database.DeleteUser(vanity)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		jsonEncoder.Encode(model.RequestError{
			Error:   true,
			Message: "Internal server error",
		})
		return
	}

	database.Set(vanity+"-gd", "ok", 3600)

	jsonEncoder.Encode(model.RequestError{
		Error:   false,
		Message: "OK",
	})
}
