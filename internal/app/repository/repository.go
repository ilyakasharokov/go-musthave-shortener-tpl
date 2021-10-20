package repository

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"ilyakasharokov/internal/app/model"
	"io"
	"log"
	"os"
)

type Repository struct {
	db              map[model.User]model.Links
	fileStoragePath string
}

type producer struct {
	file   *os.File
	writer *bufio.Writer
}

type consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func newProducer(fileName string) (*producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func newConsumer(fileName string) (*consumer, error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &consumer{
		file: f,
	}, nil
}

func (repo *Repository) AddItem(user model.User, key string, link model.Link) error {
	links := model.Links{}
	if userLinks, ok := repo.db[user]; ok {
		links = userLinks
	}
	links[key] = link
	repo.db[user] = links
	return nil
}

func (repo *Repository) GetItem(user model.User, key string) (model.Link, error) {
	return repo.db[user][key], nil
}

func (repo *Repository) GetByUser(user model.User) (model.Links, error) {
	links, ok := repo.db[user]
	if !ok {
		return links, errors.New("user not found")
	}
	return links, nil
}

func (repo *Repository) CheckExist(user model.User, key string) bool {
	links, ok := repo.db[user]
	if !ok {
		return false
	}
	_, result := links[key]
	return result
}

func New(fileStoragePath string) *Repository {
	db := make(map[model.User]model.Links)
	repo := Repository{
		db:              db,
		fileStoragePath: fileStoragePath,
	}
	repo.load()
	return &repo
}

func (repo *Repository) Flush() error {
	if repo.fileStoragePath == "" {
		return nil
	}
	// Create new producer for write links to file storage
	p, err := newProducer(repo.fileStoragePath)
	if nil != err {
		return err
	}
	// Convert to gob
	gobEncoder := gob.NewEncoder(p.writer)
	// encode
	if err := gobEncoder.Encode(repo.db); err != nil {
		return err
	}
	fmt.Println("flush")
	return p.writer.Flush()
}

// Load all links to map
func (repo *Repository) load() error {
	if repo.fileStoragePath == "" {
		return nil
	}
	cns, err := newConsumer(repo.fileStoragePath)
	if nil != err {
		return err
	}
	gobDecoder := gob.NewDecoder(cns.file)
	if err := gobDecoder.Decode(&repo.db); err != nil {
		if err != io.EOF {
			return err
		}
	}
	log.Println(repo.db)
	return nil
}
