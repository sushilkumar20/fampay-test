package service

import (
	"context"
	"fam/db"
)

type VideoQueryService interface {
	GetVideos(ctx *context.Context, page *db.Page) (*VideoMetaDataListResponse, error)
	SearchVideos(ctx *context.Context, query string, page *db.Page) (*VideoMetaDataListResponse, error)
}
