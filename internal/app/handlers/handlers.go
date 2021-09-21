package handlers

import (
	"encoding/json"
	"fmt"
	"ilyakasharokov/internal/app/base62"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/repository"
	"io"
	"net/http"
	urltool "net/url"
	"strings"
)

type URL struct {
	URL string `json:"url"`
}

const HOST = "http://localhost:8080"

func CreateShort(repo repository.RepoModel) func(w http.ResponseWriter, r *http.Request) {
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
			if !repo.CheckExist(code) {
				break
			}
		}
		err = repo.AddItem(code, link)
		if err != nil {
			http.Error(w, "Add url error", http.StatusInternalServerError)
			return
		}
		defer repo.Flush()

		result := fmt.Sprintf("%s/%s", HOST, code)
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(result))
	}
}

func APICreateShort(repo repository.RepoModel, baseURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Body read error", http.StatusBadRequest)
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
			http.Error(w, "The url is incorrect", http.StatusBadRequest)
			return
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
			if !repo.CheckExist(code) {
				break
			}
		}
		err = repo.AddItem(code, link)
		if err != nil {
			http.Error(w, "Add url error", http.StatusInternalServerError)
			return
		}
		defer repo.Flush()
		newlink := fmt.Sprintf("%s/%s", baseURL, code)
		result := struct {
			Result string `json:"result"`
		}{Result: newlink}

		body, err = json.Marshal(result)
		if err != nil {
			http.Error(w, "Response JSON error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(body)
	}
}

func GetShort(repo repository.RepoModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		id := pathSplit[1]

		entity, err := repo.GetItem(id)
		if err != nil {
			http.Error(w, "Get url error", http.StatusInternalServerError)
			return
		}

		if entity.URL == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}
