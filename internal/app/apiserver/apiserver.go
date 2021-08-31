package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"ilyakasharokov/internal/app/base62"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/repository"
	"io"
	"net/http"
	"strings"
)

const HOST = "http://localhost:8080"

type APIServer struct {
	BindAddr string
	router *mux.Router
	repo *repository.Repository
}

func New() *APIServer{
	return &APIServer{
		BindAddr: ":8080",
	}
}

func (s *APIServer) createShort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			if !s.repo.CheckExist(code) {
				break
			}
		}
		err = s.repo.AddItem(code, link)
		if err != nil {
			http.Error(w, "Add url error", http.StatusInternalServerError)
			return
		}

		result := fmt.Sprintf("%s/%s", HOST, code)
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(result))
	}
}

func (s *APIServer) getShort() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		id := pathSplit[1]

		entity, err := s.repo.GetItem(id)
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

func (s *APIServer) ConfigureRouter()  {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/", s.createShort())
	s.router.HandleFunc("/{id:[0-9a-z]+}", s.getShort())
}



func (s *APIServer) Start() error {
	s.repo = repository.New()
	s.ConfigureRouter()
	http.ListenAndServe(s.BindAddr, s.router )
	return nil
}