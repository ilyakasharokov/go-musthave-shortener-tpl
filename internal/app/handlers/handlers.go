// Обработчики HTTP сервера
package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"ilyakasharokov/internal/app/controller"
	"ilyakasharokov/internal/app/middlewares"
	"ilyakasharokov/internal/app/model"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

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
	CountURLsAndUsers(ctx context.Context) (int, int, error)
}

// CreateShort cоздает URL из тела запроса. В качестве параметра принимает репозиторий и адрес для шорта.
func CreateShort(ctrl *controller.Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
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
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}
		httpCode, shortURL, err := ctrl.CreateShort(r.Context(), url, userID)
		if err != nil {
			http.Error(w, err.Error(), httpCode)
			return
		}
		w.WriteHeader(httpCode)
		w.Write([]byte(shortURL))
	}
}

// APICreateShort запрашивает создание URL из json. В качестве параметра принимает репозиторий и адрес для шорта.
func APICreateShort(ctrl *controller.Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "body read error", http.StatusBadRequest)
			return
		}
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			userID = userIDCtx.(string)
		}
		httpCode, result, err := ctrl.APICreateShort(r.Context(), body, userID)
		if err != nil {
			http.Error(w, err.Error(), httpCode)
			return
		}
		w.WriteHeader(httpCode)
		body, _ = json.Marshal(result)
		w.Write(body)
	}
}

// GetShort получает пользовательский URL по коду. В качестве параметра принимает репозиторий.
func GetShort(repo RepoDBModel) http.HandlerFunc {
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
			return
		}
		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}

// GetUserShorts получение списка пользовательских URL. В качестве параметра принимает репозиторий.
func GetUserShorts(repo RepoDBModel) http.HandlerFunc {
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

// BunchSaveJSON загружает набор урлов в репозиторий. В качестве параметра принимает репозиторий и адрес для шорта.
func BunchSaveJSON(ctrl *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Can't read body")
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}
		var urls []model.Link
		err = json.Unmarshal(body, &urls)
		if err != nil {
			log.Err(err).Msg("Unmarshal json error")
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		httpCode, shorts, err := ctrl.BunchSaveJSON(r.Context(), urls, userID)
		if err != nil {
			http.Error(w, err.Error(), httpCode)
			return
		}
		rsp, err := json.Marshal(shorts)
		if err != nil {
			log.Err(err).Msg("Marshal links error")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(httpCode)
		w.Write(rsp)
	}
}

// Ping пингует базу данных. В качестве параметра принимает базу данных.
func Ping(db *sql.DB) http.HandlerFunc {
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

// Delete принимает множество URL в очередь на удаление.
func Delete(ctrl *controller.Controller) func(w http.ResponseWriter, r *http.Request) {
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
		userIDCtx := r.Context().Value(middlewares.UserIDCtxName)
		userID := "default"
		if userIDCtx != nil {
			// Convert interface type to user.UniqUser
			userID = userIDCtx.(string)
		}
		httpCode, err := ctrl.Delete(ids, userID)
		if err != nil {
			http.Error(w, err.Error(), httpCode)
			return
		}
		w.WriteHeader(httpCode)
	}
}

// Stats возвращает кол-во урлов и юзеров в базе. Только для доверенной подсети trustedSubnet
func Stats(repo RepoDBModel, trustedSubnet *net.IPNet) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if trustedSubnet == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		realIP := r.Header.Get("X-Real-IP")
		reqIP := net.ParseIP(realIP)
		ok := trustedSubnet.Contains(reqIP)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		users, urls, err := repo.CountURLsAndUsers(r.Context())
		if err != nil {
			http.Error(w, "sql error", http.StatusInternalServerError)
			return
		}
		response := struct {
			Urls  int `json:"urls"`
			Users int `json:"users"`
		}{
			Urls:  urls,
			Users: users,
		}
		jsn, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsn)
	}
}
