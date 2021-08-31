package repository

import (
	"ilyakasharokov/internal/app/model"
)

type Repository struct {
	db map[string]model.Link
}

func (repo *Repository) AddItem( key string, link model.Link) error{
	repo.db[key] = link
	return nil
}

func (repo *Repository) GetItem( key string) (model.Link, error){
	return repo.db[key], nil
}

func (repo *Repository) CheckExist( key string) bool{
	_, result := repo.db[key]
	return result
}

func New() *Repository{
	db := make(map[string]model.Link)
	repo := Repository{
		db: db,
	}
	return &repo
}