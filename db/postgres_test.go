package db

import (
	"context"
	"fam/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

var postgresConfig *config.PostgresConfig = &config.PostgresConfig{
	Host:         "localhost",
	UserName:     "postgres",
	Password:     "root",
	DatabaseName: "fam",
	Port:         "5432",
}
var dataBase = InitDb(postgresConfig)
var err = LoadSchema(dataBase)

func TestPostgresStore_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx     context.Context
		videoId string
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
			p := &PostgresStore{
				db: tt.fields.db,
			}
			if err := p.Delete(tt.args.ctx, tt.args.videoId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStore_DeleteBulk(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		videoIds []string
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
			p := &PostgresStore{
				db: tt.fields.db,
			}
			if err := p.DeleteBulk(tt.args.ctx, tt.args.videoIds); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBulk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStore_Get(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx     context.Context
		videoId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *VideoMetaData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgresStore{
				db: tt.fields.db,
			}
			got, err := p.Get(tt.args.ctx, tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresStore_GetPaginated(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		page *Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*VideoMetaData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgresStore{
				db: tt.fields.db,
			}
			got, err := p.GetPaginated(tt.args.ctx, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaginated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPaginated() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresStore_Insert(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx           context.Context
		videoMetaData *VideoMetaData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "insert video metadata",
			fields: fields{
				db: dataBase,
			},
			args: args{
				ctx: context.Background(),
				videoMetaData: &VideoMetaData{
					id:               uuid.New().String(),
					videoId:          uuid.New().String(),
					title:            "welcome to testing",
					description:      "enjoy your day",
					publishedAt:      time.Now().UTC().UnixMilli(),
					defaultThumbnail: "www.google.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgresStore{
				db: tt.fields.db,
			}
			if err := p.Insert(tt.args.ctx, tt.args.videoMetaData); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStore_InsertBulk(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx            context.Context
		videoMetaDatas []*VideoMetaData
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
			p := &PostgresStore{
				db: tt.fields.db,
			}
			if err := p.InsertBulk(tt.args.ctx, tt.args.videoMetaDatas); (err != nil) != tt.wantErr {
				t.Errorf("InsertBulk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStore_Search(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		query string
		page  *Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*VideoMetaData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgresStore{
				db: tt.fields.db,
			}
			got, err := p.Search(tt.args.ctx, tt.args.query, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
