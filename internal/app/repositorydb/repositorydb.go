// Модуль работы с базой данных postgres.
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

// Добавление URL в базу.
func (repo *RepositoryDB) AddItem(user model.User, key string, link model.Link, ctx context.Context) error {
	query := `
	insert into urls (id, user_id, origin_url, short_url) 
	values (default, $1, $2, $3)
	ON CONFLICT (short_url) DO NOTHING
	`
	_, err := repo.db.ExecContext(ctx, query, user, link.URL, key)
	if err != nil {
		return err
	}

	return nil
}

// Получение URL по ключу.
func (repo *RepositoryDB) GetItem(user model.User, key string, ctx context.Context) (model.Link, error) {
	query := `
		select origin_url, deleted from urls where user_id=$1 and short_url=$2
	`
	result := repo.db.QueryRowContext(ctx, query, user, key)
	link := model.Link{}
	err := result.Scan(&link.URL, &link.Deleted)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

// Удаление множества URL по id.
func (repo *RepositoryDB) RemoveItem(user model.User, id int, ctx context.Context) error {
	query := `
		update urls set deleted = true where user_id=$1 and id=$2
	`
	_, err := repo.db.ExecContext(ctx, query, user, id)
	if err != nil {
		return err
	}
	return nil
}

// Удаление множества URL по id.
func (repo *RepositoryDB) RemoveItems(user model.User, ids []int) error {
	query := `
		update urls set deleted = true where user_id=$1 and id=$2
	`
	tx, err := repo.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, id := range ids {
		_, err = tx.Exec(query, string(user), id)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Получение всех URL пользователя.
func (repo *RepositoryDB) GetByUser(user model.User, ctx context.Context) (model.Links, error) {
	query := `
		select origin_url, short_url from urls where user_id=$1
	`
	links := model.Links{}
	result, err := repo.db.QueryContext(ctx, query, user)
	if err != nil {
		return model.Links{}, err
	}
	defer func() {
		_ = result.Close()
		_ = result.Err() // or modify return value
	}()
	for result.Next() {
		link := model.Link{}
		var key string
		result.Scan(&link.URL, &key)
		links[key] = link
	}
	return links, nil
}

// Проверка существования пользовательского URL в базе.
func (repo *RepositoryDB) CheckExist(user model.User, key string) bool {
	var exist bool
	query := `
		select 1 from urls where user_id=$1 and short_url=$2
	`
	result, err := repo.db.Query(query, user, key)
	if err != nil {
		return false
	}
	defer func() {
		_ = result.Close()
		_ = result.Err() // or modify return value
	}()
	result.Scan(&exist)
	return exist
}

// Проверка существования оригинального URL в базе.
func (repo *RepositoryDB) CheckExistOrigin(user model.User, key string, ctx context.Context) model.ShortLink {
	var link = model.ShortLink{}
	query := `
		select short_url from urls where user_id=$1 and origin_url=$2
	`
	result := repo.db.QueryRowContext(ctx, query, user, key)
	err := result.Scan(&link.Short)
	if err != nil {
		return link
	}

	defer func() {
		_ = result.Err() // or modify return value
	}()
	result.Scan(&link)
	return link
}

// Сохранение множества URL.
func (repo *RepositoryDB) BunchSave(ctx context.Context, user model.User, links []model.Link) ([]model.ShortLink, error) {
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
	stmt, err := tx.PrepareContext(ctx, `
		insert into urls (id, user_id, origin_url, short_url, correlation_id) 
		values (default, $1, $2, $3, $4)
		on conflict (short_url) do nothing;
	`)
	if err != nil {
		return shorts, err
	}
	// Close statement
	defer func(stmt *sql.Stmt) {
		errr := stmt.Close()
		if errr != nil {
			fmt.Println(errr)
		}
	}(stmt)

	for _, v := range buffer {
		// Add record to transaction
		fmt.Println(v.Origin)
		if _, err = stmt.ExecContext(ctx, user, v.Origin, v.Short, v.ID); err == nil {
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

// CountURLsAndUsers возвращается кол-во юзеров и урлов
func (repo *RepositoryDB) CountURLsAndUsers(ctx context.Context) (userCount int, urlCount int, err error) {
	query := `
		SELECT COUNT(DISTINCT user_id) as count_user, COUNT(*) from urls
	`
	result := repo.db.QueryRowContext(ctx, query)
	err = result.Scan(&userCount, &urlCount)
	if err != nil {
		return 0, 0, err
	}

	defer func() {
		_ = result.Err() // or modify return value
	}()

	return userCount, urlCount, err
}

func New(db_ *sql.DB) *RepositoryDB {
	repo := RepositoryDB{
		db: db_,
	}
	return &repo
}
