package grpcshortener

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"ilyakasharokov/internal/app/controller"
	"ilyakasharokov/internal/app/model"
	"ilyakasharokov/internal/app/repositorydb"
	"ilyakasharokov/internal/app/worker"
	"ilyakasharokov/pkg/grpc/proto"
	"net"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ShortenerServer implement main methods for gRPC
type ShortenerServer struct {
	shortener.UnimplementedShortenerServer
	repo    *repositorydb.RepositoryDB
	db      *sql.DB
	wp      *worker.WorkerPool
	ctrl    *controller.Controller
	subnet  *net.IPNet
	baseURL string
}

// ResponseWriterMap it's bridge for response from main handler
type ResponseWriterMap struct {
	h    http.ResponseWriter
	head http.Header
	buf  bytes.Buffer
	code int
}

func (rw *ResponseWriterMap) Header() http.Header {
	return rw.head
}

func (rw *ResponseWriterMap) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *ResponseWriterMap) Write(data []byte) (int, error) {
	return rw.buf.Write(data)
}

// New instance for gRPC server
func New(repo *repositorydb.RepositoryDB, baseURL string, subnet *net.IPNet, db *sql.DB, wp *worker.WorkerPool, ctrl *controller.Controller) *ShortenerServer {
	return &ShortenerServer{
		shortener.UnimplementedShortenerServer{},
		repo,
		db,
		wp,
		ctrl,
		subnet,
		baseURL,
	}
}

func getUserID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	userID := "default"
	var uid []string = nil
	if ok {
		uid = md.Get("user_id")
	}
	if len(uid) > 0 {
		userID = uid[0]
	} else {
		userID = uuid.New().String()
	}
	return userID
}

func (s *ShortenerServer) CreateShort(ctx context.Context, req *shortener.URLRequest) (rsp *shortener.URLResponse, err error) {
	userID := getUserID(ctx)
	code, shortURL, err := s.ctrl.CreateShort(ctx, req.URL, userID)
	header := metadata.Pairs("user_id", userID)
	grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	rsp = new(shortener.URLResponse)
	rsp.URL = shortURL
	rsp.Code = int32(code)
	return rsp, nil
}

func (s *ShortenerServer) APICreateShort(ctx context.Context, req *shortener.URLRequest) (rsp *shortener.URLResponse, err error) {
	userID := getUserID(ctx)
	code, shortURL, err := s.ctrl.CreateShort(ctx, req.URL, userID)
	rsp = new(shortener.URLResponse)
	if err != nil {
		return nil, err
	}
	rsp.URL = shortURL
	rsp.Code = int32(code)
	return rsp, nil
}

func (s *ShortenerServer) BunchSaveJSON(ctx context.Context, req *shortener.BunchSaveRequest) (rsp *shortener.BunchSaveResponse, err error) {
	userID := getUserID(ctx)
	rLinks := req.GetLinks()
	links := []model.Link{}
	rsp = new(shortener.BunchSaveResponse)
	for _, rLink := range rLinks {
		link := model.Link{
			URL: rLink.URL,
			ID:  strconv.Itoa(int(rLink.Id)),
		}
		links = append(links, link)
	}
	_, shorts, err := s.ctrl.BunchSaveJSON(ctx, links, userID)
	if err != nil {
		return nil, err
	}
	for _, short := range shorts {
		id, _ := strconv.Atoi(short.ID)
		sLink := &shortener.BunchLink{
			URL: short.Short,
			Id:  int32(id),
		}
		rsp.Links = append(rsp.Links, sLink)
	}

	return rsp, nil
}

func (s *ShortenerServer) GetShort(ctx context.Context, req *shortener.URLRequest) (rsp *shortener.URLResponse, err error) {
	userID := getUserID(ctx)
	entity, err := s.repo.GetItem(model.User(userID), req.URL, ctx)
	rsp = new(shortener.URLResponse)
	if err != nil {
		log.Err(err).Msg("Not found")
		rsp.Code = http.StatusNotFound
		return rsp, nil
	}
	if entity.Deleted {
		log.Info().Str("id", entity.ID).Msg("Link is deleted")
		rsp.Code = http.StatusGone
		return rsp, nil
	}
	rsp.URL = entity.URL
	rsp.Code = http.StatusTemporaryRedirect
	return rsp, nil
}

func (s *ShortenerServer) GetUserShorts(ctx context.Context, _ *shortener.Empty) (rsp *shortener.GetUserShortsResponse, err error) {
	userID := getUserID(ctx)
	links, err := s.repo.GetByUser(model.User(userID), ctx)
	rsp = new(shortener.GetUserShortsResponse)
	if err != nil {
		log.Err(err).Msg("marshal links error")
		rsp.Code = http.StatusInternalServerError
		return rsp, nil
	}
	if len(links) == 0 {
		log.Err(err).Str("user", userID).Msg("no links")
		rsp.Code = http.StatusNoContent
		return rsp, nil
	}
	rsp.Code = http.StatusOK
	for _, link := range links {
		id, _ := strconv.Atoi(link.ID)
		rsp.Links = append(rsp.Links, &shortener.BunchLink{
			URL: link.URL,
			Id:  int32(id),
		})
	}
	return rsp, nil
}
func (s *ShortenerServer) Ping(ctx context.Context, _ *shortener.Empty) (rsp *shortener.CodeResponse, err error) {
	err = s.db.PingContext(ctx)
	rsp = new(shortener.CodeResponse)
	if err != nil {
		fmt.Println(err)
		rsp.Code = http.StatusInternalServerError
		return rsp, nil
	}
	rsp.Code = http.StatusOK
	return rsp, nil
}

func (s *ShortenerServer) Delete(ctx context.Context, req *shortener.DeleteRequest) (rsp *shortener.CodeResponse, err error) {
	userID := getUserID(ctx)
	idsInt := []int{}
	for _, id := range req.Id {
		idsInt = append(idsInt, int(id))
	}
	rsp = new(shortener.CodeResponse)
	code, err := s.ctrl.Delete(idsInt, userID)
	if err != nil {
		return nil, err
	}
	rsp.Code = int32(code)
	return rsp, err
}

func (s *ShortenerServer) Stats(ctx context.Context, _ *shortener.Empty) (rsp *shortener.StatsResponse, err error) {
	rsp = new(shortener.StatsResponse)
	if s.subnet == nil {
		rsp.Code = http.StatusForbidden
		return rsp, nil
	}
	var realIP string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		realIPArr := md.Get("X-Real-IP")
		if len(realIPArr) > 0 {
			realIP = realIPArr[0]
		}
	}
	if realIP == "" {
		rsp.Code = http.StatusForbidden
		return rsp, nil
	}
	reqIP := net.ParseIP(realIP)
	ok = s.subnet.Contains(reqIP)
	if !ok {
		rsp.Code = http.StatusForbidden
		return rsp, nil
	}
	users, urls, err := s.repo.CountURLsAndUsers(ctx)
	if err != nil {
		rsp.Code = http.StatusInternalServerError
		return rsp, nil
	}
	rsp.Code = http.StatusOK
	rsp.Users = int32(users)
	rsp.URLs = int32(urls)
	return rsp, nil
}
