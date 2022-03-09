package repositorydb

import (
	"context"
	"database/sql"
	"ilyakasharokov/internal/app/model"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		db_ *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *RepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_AddItem(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		key  string
		link model.Link
		ctx  context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			if err := repo.AddItem(tt.args.user, tt.args.key, tt.args.link, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepositoryDB_BunchSave(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		user  model.User
		links []model.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.ShortLink
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			got, err := repo.BunchSave(tt.args.ctx, tt.args.user, tt.args.links)
			if (err != nil) != tt.wantErr {
				t.Errorf("BunchSave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BunchSave() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_CheckExist(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		key  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			if got := repo.CheckExist(tt.args.user, tt.args.key); got != tt.want {
				t.Errorf("CheckExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_CheckExistOrigin(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		key  string
		ctx  context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.ShortLink
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			if got := repo.CheckExistOrigin(tt.args.user, tt.args.key, tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckExistOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_GetByUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		ctx  context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Links
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			got, err := repo.GetByUser(tt.args.user, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_GetItem(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		key  string
		ctx  context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Link
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			got, err := repo.GetItem(tt.args.user, tt.args.key, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryDB_RemoveItem(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		id   int
		ctx  context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			if err := repo.RemoveItem(tt.args.user, tt.args.id, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepositoryDB_RemoveItems(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user model.User
		ids  []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &RepositoryDB{
				db: tt.fields.db,
			}
			if err := repo.RemoveItems(tt.args.user, tt.args.ids); (err != nil) != tt.wantErr {
				t.Errorf("RemoveItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
