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
func (c *Controller) CreateShort(ctx context.Context, url string, userID string) (err error, httpCode int, shortURL string) {
	_, err = urltool.ParseRequestURI(url)
	if err != nil {
		return errors.New("The url is incorrect"), http.StatusBadRequest, ""
	}
	var code string
	for {
		code, err = base62.Decode(url)
		if err != nil {
			return errors.New("URL decode error"), http.StatusInternalServerError, ""
		}
		if !c.repo.CheckExist(model.User(userID), code) {
			break
		}
	}
	link, err := c.repo.GetItem(model.User(userID), code, ctx)
	shortURL = fmt.Sprintf("%s/%s", c.baseURL, code)
	if err == nil {
		return errors.New("Already exist"), http.StatusConflict, shortURL
	}
	link.URL = url
	err = c.repo.AddItem(model.User(userID), code, link, ctx)
	if err != nil {
		return errors.New("Add url error"), http.StatusInternalServerError, ""
	}
	return nil, http.StatusCreated, shortURL
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

func (c *Controller) APICreateShort(ctx context.Context, data []byte, userID string) (err error, httpCode int, result model.Result) {
	url, err := c.ParseJSON(data)
	if err != nil {
		return err, http.StatusBadRequest, result
	}
	link := model.Link{
		URL: url.URL,
	}
	var code string
	for {
		code, err = base62.Decode(link.URL)
		if err != nil {
			return errors.New("URL decode error"), http.StatusInternalServerError, result
		}
		if !c.repo.CheckExist(model.User(userID), code) {
			break
		}
	}

	_, err = c.repo.GetItem(model.User(userID), code, ctx)
	newlink := fmt.Sprintf("%s/%s", c.baseURL, code)
	result = model.Result{Result: newlink}
	if err == nil {
		return nil, http.StatusConflict, result
	}

	err = c.repo.AddItem(model.User(userID), code, link, ctx)
	if err != nil {
		return errors.New("Add url error"), http.StatusInternalServerError, result
	}
	return nil, http.StatusCreated, result
}

func (c *Controller) BunchSaveJSON(ctx context.Context, links []model.Link, userID string) (err error, httpCode int, shorts []model.ShortLink) {
	shorts, err = c.repo.BunchSave(ctx, model.User(userID), links)
	if err != nil {
		log.Err(err).Msg("Can't save links")
		return errors.New("can't save"), http.StatusBadRequest, nil
	}
	// Prepare results
	for k := range shorts {
		shorts[k].Short = fmt.Sprintf("%s/%s", c.baseURL, shorts[k].Short)
	}
	return nil, http.StatusCreated, shorts
}

func (c *Controller) GetShort() {

}

func (c *Controller) Delete(ids []int, userID string) (err error, httpCode int) {
	if len(ids) == 0 {
		return errors.New("No ids"), http.StatusBadRequest
		return
	}
	bf := func(_ context.Context) error {
		c.repo.RemoveItems(model.User(userID), ids)
		return nil
	}
	c.wp.Push(bf)
	return nil, http.StatusAccepted
}

/*
r.Post("/api/shorten/batch", handlers.BunchSaveJSON(repo, baseURL))
r.Get("/{id:[0-9a-zA-z]+}", handlers.GetShort(repo))
r.Get("/user/urls", handlers.GetUserShorts(repo))
r.Get("/ping", handlers.Ping(database))
r.Delete("/api/user/urls", handlers.Delete(repo, wp))
r.Get("/api/internal/stats", handlers.Stats(repo, trustedSubnet))*/
