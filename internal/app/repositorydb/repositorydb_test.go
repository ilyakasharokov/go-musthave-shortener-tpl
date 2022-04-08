package repositorydb

import (
	"context"
	"database/sql"
	"ilyakasharokov/internal/app/model"
	"testing"
)

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
