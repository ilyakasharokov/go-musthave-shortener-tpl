package repositorydb

import (
	"context"
	"database/sql"
	"fmt"
	helpers "ilyakasharokov/internal/app/encryptor"
	"ilyakasharokov/internal/app/model"
)

type RepositoryDB struct {
	db *sql.DB
}

func (repo *RepositoryDB) AddItem(user model.User, key string, link model.Link) error {
	sql := `
	insert into urls (id, user_id, origin_url, short_url) 
	values (default, $1, $2, $3)
	ON CONFLICT (short_url) DO NOTHING
	`
	_, err := repo.db.ExecContext(context.Background(), sql, user, link.URL, key)
	if err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryDB) GetItem(user model.User, key string) (model.Link, error) {
	sql := `
		select origin_url from urls where user_id=$1 and short_url=$2
	`
	result := repo.db.QueryRowContext(context.Background(), sql, user, key)
	link := model.Link{}
	err := result.Scan(&link.URL)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

func (repo *RepositoryDB) GetByUser(user model.User) (model.Links, error) {
	sql := `
		select origin_url, short_url from urls where user_id=$1
	`
	links := model.Links{}
	result, err := repo.db.QueryContext(context.Background(), sql, user)
	if err != nil {
		return model.Links{}, err
	}
	for result.Next() {
		link := model.Link{}
		var key string
		result.Scan(&link.URL, &key)
		links[key] = link
	}
	return links, nil
}

func (repo *RepositoryDB) CheckExist(user model.User, key string) bool {
	var exist bool
	sql := `
		select 1 from urls where user_id=$1 and short_url=$2
	`
	result, err := repo.db.Query(sql, user, key)
	if err != nil {
		return false
	}
	result.Scan(&exist)
	return exist
}

func (repo *RepositoryDB) BunchSave(links []model.Link) ([]model.ShortLink, error) {
	// Generate shorts
	type temp struct {
		ID,
		Origin,
		Short string
	}

	var buffer []temp
	for _, v := range links {
		var t = temp{
			ID:     v.ID,
			Origin: v.URL,
			Short:  helpers.RandomString(10),
		}
		buffer = append(buffer, t)
	}
	dbd := repo.db
	var shorts []model.ShortLink

	// Delete old records for tests
	_, _ = dbd.Exec("truncate table urls;")

	// Start transaction
	tx, err := dbd.Begin()
	if err != nil {
		return shorts, err
	}
	// Rollback handler
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)
	// Prepare statement
	stmt, err := tx.PrepareContext(context.Background(), `
		insert into urls (id, user_id, origin_url, short_url, correlation_id) 
		values (default, $1, $2, $3, $4)
		on conflict (short_url) do nothing;
	`)
	if err != nil {
		return shorts, err
	}
	// Close statement
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(stmt)

	for _, v := range buffer {
		// Add record to transaction
		if _, err = stmt.ExecContext(context.Background(), "e210091c-3196-11ec-b01c-3e22fb9798bf", v.Origin, v.Short, v.ID); err == nil {
			shorts = append(shorts, model.ShortLink{
				Short: v.Short,
				ID:    v.ID,
			})
		} else {
			return shorts, err
		}
	}
	// шаг 4 — сохраняем изменения
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return shorts, nil
}

func New(db_ *sql.DB) *RepositoryDB {
	repo := RepositoryDB{
		db: db_,
	}
	return &repo
}
