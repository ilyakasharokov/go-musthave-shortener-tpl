package repository

import (
	"ilyakasharokov/internal/app/model"
	"reflect"
	"testing"
)

const testUser = model.User("default")
const testCode = "1692759882237307797"

func TestRepository_AddItem(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
	}
	type args struct {
		user model.User
		key  string
		link model.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "good payload",
			args: args{
				user: testUser,
				key:  testCode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			repo.db = make(map[model.User]model.Links)
			if err := repo.AddItem(tt.args.user, tt.args.key, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_CheckExist(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
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
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			if got := repo.CheckExist(tt.args.user, tt.args.key); got != tt.want {
				t.Errorf("CheckExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Flush(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			if err := repo.Flush(); (err != nil) != tt.wantErr {
				t.Errorf("Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetByUser(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
	}
	type args struct {
		user model.User
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
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			got, err := repo.GetByUser(tt.args.user)
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

func TestRepository_GetItem(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
	}
	type args struct {
		user model.User
		key  string
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
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			got, err := repo.GetItem(tt.args.user, tt.args.key)
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

func TestRepository_load(t *testing.T) {
	type fields struct {
		db              map[model.User]model.Links
		fileStoragePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				db:              tt.fields.db,
				fileStoragePath: tt.fields.fileStoragePath,
			}
			if err := repo.load(); (err != nil) != tt.wantErr {
				t.Errorf("load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newConsumer(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *consumer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newConsumer(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("newConsumer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConsumer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newProducer(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *producer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newProducer(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("newProducer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newProducer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
