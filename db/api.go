package db

import (
	"context"
)

type DBClinet interface {
	Insert(ctx context.Context, videoMetaData *VideoMetaData) error
	InsertBulk(ctx context.Context, videoMetaData []*VideoMetaData) error
	Delete(ctx context.Context, videoId string) error
	DeleteBulk(ctx context.Context, videoIds []string) error
	Get(ctx context.Context, videoId string) (*VideoMetaData, error)
	GetPaginated(ctx context.Context, page *Page) ([]*VideoMetaData, error)
	Search(ctx context.Context, searchQuery string, page *Page) ([]*VideoMetaData, error)
}
