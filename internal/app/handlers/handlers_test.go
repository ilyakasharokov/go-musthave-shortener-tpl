package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/mocks"
	"ilyakasharokov/internal/app/model"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testURL = "https://yandex.ru"
const testUser = model.User("default")
const testCode = "1692759882237307797"

var cfg = configuration.Config{
	BaseURL:         "http://example.com",
	FileStoragePath: "",
}

func TestCreateShort(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name    string
		payload string
		want    want
	}{
		{
			name:    "#1 post request test good payload",
			payload: testURL,
			want: want{
				code:        http.StatusCreated,
				contentType: "",
			},
		},
		{
			name:    "#2 post request test empty payload",
			payload: "",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "#2 post request test not an url",
			payload: "asdfasfsa",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	repo := new(mocks.RepoDBModel)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.On("CheckExist", model.User(testUser), testCode).Return(false)
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.payload))
			repo.On("GetItem", model.User(testUser), testCode, request.Context()).Return(model.Link{}, errors.New("not found"))
			repo.On("AddItem", model.User(testUser), testCode, model.Link{URL: ""}, request.Context()).Return(nil)
			repo.On("AddItem", model.User(testUser), testCode, model.Link{URL: tt.payload}, request.Context()).Return(nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(CreateShort(repo, cfg.BaseURL))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func ExampleCreateShort() {
	repo := new(mocks.RepoDBModel)
	r := chi.NewRouter()
	r.Post("/", CreateShort(repo, "http://example.com"))
}

func TestGetShort(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "#1 get request test",
			path: testCode,
			want: want{
				code:        http.StatusTemporaryRedirect,
				contentType: "application/text",
			},
		},
		{
			name: "#2 get request test",
			path: "_",
			want: want{
				code:        http.StatusNotFound,
				contentType: "application/text",
			},
		},
	}

	repo := new(mocks.RepoDBModel)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", tt.path), nil)
			repo.On("GetItem", model.User(testUser), "_", request.Context()).Return(model.Link{}, errors.New("Not found"))
			repo.On("GetItem", model.User(testUser), testCode, request.Context()).Return(model.Link{URL: testURL}, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(GetShort(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)
		})
	}
}

func TestAPICreateShort(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name          string
		payload       string
		addItemResult error
		want          want
	}{
		{
			name:          "#1 post request test good payload",
			payload:       testURL,
			addItemResult: nil,
			want: want{
				code:        http.StatusCreated,
				contentType: "",
			},
		},
		{
			name:          "#2 post request test empty payload",
			payload:       "",
			addItemResult: errors.New("add url error"),
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:          "#2 post request test not an url",
			payload:       "asdfasfsa",
			addItemResult: errors.New("add url error"),
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	repo := new(mocks.RepoDBModel)
	repo.On("CheckExist", model.User(testUser), testCode).Return(false)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.payload))
			repo.On("GetItem", model.User(testUser), testCode, request.Context()).Return(model.Link{}, errors.New("not found"))
			repo.On("AddItem", model.User(testUser), testCode, model.Link{URL: tt.payload}, request.Context()).Return(tt.addItemResult)
			repo.On("AddItem", model.User(testUser), testCode, model.Link{URL: ""}, request.Context()).Return(tt.addItemResult)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(CreateShort(repo, cfg.BaseURL))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func TestBunchSaveJSON(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name          string
		payload       string
		addItemResult error
		want          want
	}{
		{
			name: "#1 post request test good payload",
			payload: `[
				{
					"correlation_id":"1",
					"original_url":"` + testURL + `"
				}]`,
			addItemResult: nil,
			want: want{
				code: http.StatusCreated,
			},
		},
	}

	repo := new(mocks.RepoDBModel)
	repo.On("CheckExist", model.User(testUser), testCode).Return(false)
	repo.On("BunchSave", context.Background(), model.User(testUser), []model.Link{{ID: "1", URL: testURL}}).Return([]model.ShortLink{{ID: "1", Short: testCode}}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(BunchSaveJSON(repo, cfg.BaseURL))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)
		})
	}
}

func ExampleBunchSaveJSON() {
	repo := new(mocks.RepoDBModel)
	r := chi.NewRouter()
	r.Post("/api/shorten/batch", BunchSaveJSON(repo, "http://example.com"))

	/*
		curl --location --request POST 'http://localhost:8080/api/shorten/batch' \
		--header 'Content-Type: application/json' \
		--data-raw '[
			{
				"correlation_id":"1",
				"original_url":"http://yandex.ru"
			},
			{
				"correlation_id":"2",
				"original_url":"http://google.com"
			}
		]'

	*/
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkCreateShort(b *testing.B) {
	repo := new(mocks.RepoDBModel)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		url := RandStringBytes(10)
		b.StartTimer()
		CreateShort(repo, url)
	}
}
