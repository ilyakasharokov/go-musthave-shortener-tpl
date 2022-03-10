// Обработчики HTTP сервера
package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"ilyakasharokov/internal/app/base62"
	"ilyakasharokov/internal/app/middlewares"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/worker"
	"io"
	"io/ioutil"
	"net/http"
	urltool "net/url"
	"strings"

	"github.com/rs/zerolog/log"
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
	AddItem(model.User, string, model.Link, context.Context) error
	GetItem(model.User, string, context.Context) (model.Link, error)
	CheckExist(model.User, string) bool
	GetByUser(model.User, context.Context) (model.Links, error)
	BunchSave(context.Context, model.User, []model.Link) ([]model.ShortLink, error)
	RemoveItems(model.User, []int) error
}

// Запрос на создание URL из тела запроса.
func CreateShort(repo RepoDBModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
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
		link, err = repo.GetItem(model.User(userID), code, r.Context())
		result := fmt.Sprintf("%s/%s", baseURL, code)
		if err == nil {
			http.Error(w, "Already exist", http.StatusConflict)
			w.Header().Add("Content-type", "text/plain; charset=utf-8")
			w.Write([]byte(result))
			return
		}

		err = repo.AddItem(model.User(userID), code, link, r.Context())
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

// Запрос на создание URL из json.
func APICreateShort(repo RepoDBModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
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

		_, err = repo.GetItem(model.User(userID), code, r.Context())
		newlink := fmt.Sprintf("%s/%s", baseURL, code)
		result := struct {
			Result string `json:"result"`
		}{Result: newlink}
		if err == nil {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusConflict)
			body, err = json.Marshal(result)
			if err != nil {
				http.Error(w, "response JSON error", http.StatusInternalServerError)
				return
			}
			w.Write(body)
			return
		}

		err = repo.AddItem(model.User(userID), code, link, r.Context())
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

		w.Write(body)
	}
}

// Получение пользовательского URL по коду.
func GetShort(repo RepoDBModel) func(w http.ResponseWriter, r *http.Request) {
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

		entity, err := repo.GetItem(model.User(userID), id, r.Context())

		if err != nil {
			log.Err(err).Msg("Not found")
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if entity.Deleted {
			log.Info().Str("id", entity.ID).Msg("Link is deleted")
			http.Error(w, "Deleted", http.StatusGone)
		}
		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}

// Получение списка пользовательских URL.
func GetUserShorts(repo RepoDBModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		links, err := repo.GetByUser(model.User(userID), r.Context())
		if err != nil {
			log.Err(err).Str("user", userID).Msg("No links")
			http.Error(w, "no content", http.StatusNoContent)
			return
		}

		body, err := links.MarshalJSON()
		if err != nil {
			log.Err(err).Msg("Marshal links error")
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
			log.Err(err).Msg("Can't read body")
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}
		// Get url from json data
		var urls []model.Link
		err = json.Unmarshal(body, &urls)
		if err != nil {
			log.Err(err).Msg("Unmarshal json error")
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		shorts, err := repo.BunchSave(r.Context(), model.User(userID), urls)
		if err != nil {
			log.Err(err).Msg("Can't save links")
			http.Error(w, "can't save", http.StatusBadRequest)
			return
		}
		// Prepare results
		for k := range shorts {
			shorts[k].Short = fmt.Sprintf("%s/%s", baseURL, shorts[k].Short)
		}

		body, err = json.Marshal(shorts)
		if err != nil {
			log.Err(err).Msg("Marshal links error")
			http.Error(w, "Marshal error", http.StatusBadRequest)
		}

		// Prepare response
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(body)
		if err != nil {
			log.Err(err).Msg("Body write error")
			http.Error(w, "Body write error", http.StatusBadRequest)
		}

	}
}

// Пинг базы данных.
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

// Запрос на удаление множества URL.
func Delete(repo RepoDBModel, workerPool *worker.WorkerPool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "body read error", http.StatusBadRequest)
			return
		}
		var ids []int
		err = json.Unmarshal(body, &ids)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "JSON is incorrect", http.StatusBadRequest)
			return
		}
		if len(ids) == 0 {
			http.Error(w, "No ids", http.StatusBadRequest)
			return
		}
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}

		bf := func(_ context.Context) error {
			repo.RemoveItems(model.User(userID), ids)
			return nil
		}

		workerPool.Push(bf)
		w.WriteHeader(http.StatusAccepted)

	}
}
