package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testUrl = "https://yandex.ru"

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
			payload: testUrl,
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain; charset=utf-8",
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
	repo := repository.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(CreateShort(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, res.Header.Get("Content-Type"), tt.want.contentType)
		})
	}
}

func TestGetShort(t *testing.T) {
	const testCode = "1692759882237307797"
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

	repo := repository.New()
	repo.AddItem(testCode, model.Link{URL: testUrl})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", tt.path), nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(GetShort(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)
		})
	}
}
