package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"ilyakasharokov/internal/app/base62"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/worker"
	"net"
	"net/http"
	urltool "net/url"

	"github.com/rs/zerolog/log"
)

type RepoDBModel interface {
	AddItem(model.User, string, model.Link, context.Context) error
	GetItem(model.User, string, context.Context) (model.Link, error)
	CheckExist(model.User, string) bool
	GetByUser(model.User, context.Context) (model.Links, error)
	BunchSave(context.Context, model.User, []model.Link) ([]model.ShortLink, error)
	RemoveItems(model.User, []int) error
	CountURLsAndUsers(ctx context.Context) (int, int, error)
}

type URL struct {
	URL string `json:"url"`
}

type Controller struct {
	repo    RepoDBModel
	baseURL string
	db      *sql.DB
	wp      *worker.WorkerPool
	subnet  *net.IPNet
}

func NewController(repo RepoDBModel, baseURL string, db *sql.DB, wp *worker.WorkerPool, subnet *net.IPNet) *Controller {
	return &Controller{
		repo,
		baseURL,
		db,
		wp,
		subnet,
	}
}

// CreateShort cоздает URL из параметра url для юзера userID. Возвращает ошибку и shortURL
func (c *Controller) CreateShort(ctx context.Context, url string, userID string) (httpCode int, shortURL string, err error) {
	_, err = urltool.ParseRequestURI(url)
	if err != nil {
		return http.StatusBadRequest, "", errors.New("the url is incorrect")
	}
	var code string
	for {
		code, err = base62.Decode(url)
		if err != nil {
			return http.StatusInternalServerError, "", errors.New("URL decode error")
		}
		if !c.repo.CheckExist(model.User(userID), code) {
			break
		}
	}
	link, err := c.repo.GetItem(model.User(userID), code, ctx)
	shortURL = fmt.Sprintf("%s/%s", c.baseURL, code)
	if err == nil {
		return http.StatusConflict, shortURL, errors.New("already exist")
	}
	link.URL = url
	err = c.repo.AddItem(model.User(userID), code, link, ctx)
	if err != nil {
		return http.StatusInternalServerError, "", errors.New("add url error")
	}
	return http.StatusCreated, shortURL, nil
}

func (c *Controller) ParseJSON(data []byte) (URL, error) {
	url := URL{}
	err := json.Unmarshal(data, &url)
	if err != nil {
		return url, errors.New("JSON is incorrect")
	}
	if url.URL == "" {
		return url, errors.New("URL is empty")
	}
	_, err = urltool.ParseRequestURI(url.URL)
	if err != nil {
		return url, errors.New("the url is incorrect")
	}
	return url, nil
}

func (c *Controller) APICreateShort(ctx context.Context, data []byte, userID string) (httpCode int, result model.Result, err error) {
	url, err := c.ParseJSON(data)
	if err != nil {
		return http.StatusBadRequest, result, err
	}
	link := model.Link{
		URL: url.URL,
	}
	var code string
	for {
		code, err = base62.Decode(link.URL)
		if err != nil {
			return http.StatusInternalServerError, result, errors.New("URL decode error")
		}
		if !c.repo.CheckExist(model.User(userID), code) {
			break
		}
	}

	_, err = c.repo.GetItem(model.User(userID), code, ctx)
	newlink := fmt.Sprintf("%s/%s", c.baseURL, code)
	result = model.Result{Result: newlink}
	if err == nil {
		return http.StatusConflict, result, nil
	}

	err = c.repo.AddItem(model.User(userID), code, link, ctx)
	if err != nil {
		return http.StatusInternalServerError, result, errors.New("add url error")
	}
	return http.StatusCreated, result, nil
}

func (c *Controller) BunchSaveJSON(ctx context.Context, links []model.Link, userID string) (httpCode int, shorts []model.ShortLink, err error) {
	shorts, err = c.repo.BunchSave(ctx, model.User(userID), links)
	if err != nil {
		log.Err(err).Msg("Can't save links")
		return http.StatusBadRequest, nil, errors.New("can't save")
	}
	// Prepare results
	for k := range shorts {
		shorts[k].Short = fmt.Sprintf("%s/%s", c.baseURL, shorts[k].Short)
	}
	return http.StatusCreated, shorts, nil
}

func (c *Controller) GetShort() {

}

func (c *Controller) Delete(ids []int, userID string) (httpCode int, err error) {
	if len(ids) == 0 {
		return http.StatusBadRequest, errors.New("no ids")
	}
	bf := func(_ context.Context) error {
		c.repo.RemoveItems(model.User(userID), ids)
		return nil
	}
	c.wp.Push(bf)
	return http.StatusAccepted, nil
}

/*
r.Post("/api/shorten/batch", handlers.BunchSaveJSON(repo, baseURL))
r.Get("/{id:[0-9a-zA-z]+}", handlers.GetShort(repo))
r.Get("/user/urls", handlers.GetUserShorts(repo))
r.Get("/ping", handlers.Ping(database))
r.Delete("/api/user/urls", handlers.Delete(repo, wp))
r.Get("/api/internal/stats", handlers.Stats(repo, trustedSubnet))*/
