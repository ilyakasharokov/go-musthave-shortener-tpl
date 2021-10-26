package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"ilyakasharokov/internal/app/base62"
	"ilyakasharokov/internal/app/middlewares"
	"ilyakasharokov/internal/app/model"
	"io"
	"io/ioutil"
	"net/http"
	urltool "net/url"
	"strings"
)

type URL struct {
	URL string `json:"url"`
}

type RepoModel interface {
	AddItem(model.User, string, model.Link) error
	GetItem(model.User, string) (model.Link, error)
	CheckExist(model.User, string) bool
	GetByUser(model.User) (model.Links, error)
}

type RepoDBModel interface {
	AddItem(model.User, string, model.Link) error
	GetItem(model.User, string) (model.Link, error)
	CheckExist(model.User, string) bool
	GetByUser(model.User) (model.Links, error)
	BunchSave([]model.Link) ([]model.ShortLink, error)
}

func CreateShort(repo RepoModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Body read error", http.StatusBadRequest)
			return
		}
		url := string(body)
		if url == "" {
			http.Error(w, "Body is empty", http.StatusBadRequest)
			return
		}

		_, err = urltool.ParseRequestURI(url)

		if err != nil {
			http.Error(w, "The url is incorrect", http.StatusBadRequest)
			return
		}

		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		link := model.Link{
			URL: url,
		}

		var code string
		for {
			code, err = base62.Decode(link.URL)
			if err != nil {
				http.Error(w, "URL decode error", http.StatusInternalServerError)
				return
			}
			if !repo.CheckExist(model.User(userID), code) {
				break
			}
		}
		link, err = repo.GetItem(model.User(userID), code)
		result := fmt.Sprintf("%s/%s", baseURL, code)
		if err == nil {
			http.Error(w, "Already exist", http.StatusConflict)
			w.Header().Add("Content-type", "text/plain; charset=utf-8")
			w.Write([]byte(result))
			return
		}

		err = repo.AddItem(model.User(userID), code, link)
		if err != nil {
			http.Error(w, "Add url error", http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.Write([]byte(result))
	}
}

func APICreateShort(repo RepoModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "body read error", http.StatusBadRequest)
			return
		}
		url := URL{}
		err = json.Unmarshal(body, &url)
		if err != nil {
			http.Error(w, "JSON is incorrect", http.StatusBadRequest)
			return
		}
		if url.URL == "" {
			http.Error(w, "URL is empty", http.StatusBadRequest)
			return
		}

		_, err = urltool.ParseRequestURI(url.URL)

		if err != nil {
			http.Error(w, "the url is incorrect", http.StatusBadRequest)
			return
		}

		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		link := model.Link{
			URL: url.URL,
		}

		var code string
		for {
			code, err = base62.Decode(link.URL)
			if err != nil {
				http.Error(w, "URL decode error", http.StatusInternalServerError)
				return
			}
			if !repo.CheckExist(model.User(userID), code) {
				break
			}
		}

		link, err = repo.GetItem(model.User(userID), code)
		newlink := fmt.Sprintf("%s/%s", baseURL, code)
		result := struct {
			Result string `json:"result"`
		}{Result: newlink}
		if err == nil {
			http.Error(w, "Already exist", http.StatusConflict)
			body, err = json.Marshal(result)
			if err != nil {
				http.Error(w, "response JSON error", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Write(body)
			return
		}

		err = repo.AddItem(model.User(userID), code, link)
		if err != nil {
			http.Error(w, "Add url error", http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}

		body, err = json.Marshal(result)
		if err != nil {
			http.Error(w, "response JSON error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(body)
	}
}

func GetShort(repo RepoModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		id := pathSplit[1]

		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		entity, err := repo.GetItem(model.User(userID), id)

		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}

func GetUserShorts(repo RepoModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		links, err := repo.GetByUser(model.User(userID))
		if err != nil {
			http.Error(w, "no content", http.StatusNoContent)
			return
		}

		body, err := links.MarshalJSON()
		if err != nil {
			http.Error(w, "json error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func bodyFromJSON(w *http.ResponseWriter, r *http.Request) ([]byte, error) {
	var body []byte
	if r.Body == http.NoBody {
		http.Error(*w, "no content", http.StatusBadRequest)
		return body, errors.New("bad request")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(*w, "unknown url", http.StatusBadRequest)
		return body, errors.New("unknown url")
	}
	return body, nil
}

// BunchSaveJSON save data and return from mass
func BunchSaveJSON(repo RepoDBModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := bodyFromJSON(&w, r)
		if err != nil {
			http.Error(w, "unknown url", http.StatusBadRequest)
			return
		}
		// Get url from json data
		var urls []model.Link
		err = json.Unmarshal(body, &urls)
		if err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		shorts, err := repo.BunchSave(urls)
		if err != nil {
			http.Error(w, "can't save", http.StatusBadRequest)
			return
		}
		// Prepare results
		for k := range shorts {
			shorts[k].Short = fmt.Sprintf("%s/%s", baseURL, shorts[k].Short)
		}

		body, err = json.Marshal(shorts)
		if err == nil {
			// Prepare response
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusCreated)
			_, err = w.Write(body)
			if err == nil {
				return
			}
		}
		http.Error(w, "DB error", http.StatusBadRequest)
	}
}

func Ping(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := db.PingContext(r.Context())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
