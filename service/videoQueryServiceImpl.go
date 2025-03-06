package service

import (
	"context"
	"fam/db"
	"time"
)

type VideoQueryServiceImpl struct {
	dbClient db.DBClinet
}

func NewVideoQueryServiceImpl(dbClient db.DBClinet) *VideoQueryServiceImpl {
	return &VideoQueryServiceImpl{dbClient: dbClient}
}

func (obj *VideoQueryServiceImpl) SearchVideos(ctx *context.Context, searchQuery string, page *db.Page) (*VideoMetaDataListResponse, error) {
	videos, err := obj.dbClient.Search(*ctx, searchQuery, page)
	if err != nil {
		return nil, err
	}

	return obj.parse(videos), nil
}

func (obj *VideoQueryServiceImpl) GetVideos(ctx *context.Context, page *db.Page) (*VideoMetaDataListResponse, error) {
	videos, err := obj.dbClient.GetPaginated(*ctx, page)

	if err != nil {
		return nil, err
	}

	return obj.parse(videos), nil
}

func (obj *VideoQueryServiceImpl) parse(videos []*db.VideoMetaData) *VideoMetaDataListResponse {
	var videoResponseData []VideoResponseData

	for _, video := range videos {

		publishedTime := time.UnixMilli(video.PublishedAt)
		rfcTime := publishedTime.Format(time.RFC3339)
		videoResponseData = append(videoResponseData, VideoResponseData{
			VideoId:          video.VideoId,
			Title:            video.Title,
			Description:      video.Description,
			PublishedAt:      rfcTime,
			DefaultThumbnail: video.DefaultThumbnail,
		})
	}

	return &VideoMetaDataListResponse{
		Videos: videoResponseData,
	}
}
