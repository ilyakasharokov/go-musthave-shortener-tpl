package repository

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/model"
	"io"
	"log"
	"os"
)

type RepoModel interface {
	AddItem(string, model.Link) error
	GetItem(string) (model.Link, error)
	CheckExist(string) bool
	Flush() error
	Load() error
}

type Repository struct {
	db     map[string]model.Link
	config configuration.Config
}

type Producer interface {
	WriteEvent(event *Repository)
	Close() error
}

type Consumer interface {
	ReadEvent() (*Repository, error)
	Close() error
}

type producer struct {
	file   *os.File
	writer *bufio.Writer
}

type consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewProducer(fileName string) (*producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func NewConsumer(fileName string) (*consumer, error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &consumer{
		file: f,
	}, nil
}

func (p *producer) Close() error {
	return p.file.Close()
}

func (c *consumer) Close() error {
	return c.file.Close()
}

func (repo *Repository) AddItem(key string, link model.Link) error {
	repo.db[key] = link
	return nil
}

func (repo *Repository) GetItem(key string) (model.Link, error) {
	return repo.db[key], nil
}

func (repo *Repository) CheckExist(key string) bool {
	_, result := repo.db[key]
	return result
}

func New(cfg configuration.Config) *Repository {
	db := make(map[string]model.Link)
	repo := Repository{
		db:     db,
		config: cfg,
	}
	repo.Load()
	return &repo
}

func (repo *Repository) Flush() error {
	if repo.config.FileStoragePath == "" {
		return nil
	}
	// Create new producer for write links to file storage
	p, err := NewProducer(repo.config.FileStoragePath)
	if nil != err {
		return err
	}
	// Convert to gob
	gobEncoder := gob.NewEncoder(p.writer)
	// encode
	if err := gobEncoder.Encode(repo.db); err != nil {
		panic(err)
	}
	fmt.Println("flush")
	return p.writer.Flush()
}

// Load all links to map
func (repo *Repository) Load() error {
	if repo.config.FileStoragePath == "" {
		return nil
	}
	cns, err := NewConsumer(repo.config.FileStoragePath)
	if nil != err {
		return err
	}
	gobDecoder := gob.NewDecoder(cns.file)
	if err := gobDecoder.Decode(&repo.db); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	log.Println(repo.db)
	return nil
}
