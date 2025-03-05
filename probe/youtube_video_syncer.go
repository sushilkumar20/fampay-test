package probe

import (
	"context"
	"fam/db"
	"fam/manager"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type YouTubeVideoSyncer struct {
	db.DBClinet
	manager.YouTubeClient
}

// TODO : check logic for time
func (y *YouTubeVideoSyncer) Execute() error {
	ctx := context.Background()
	lastFetchedTime := time.Now().UTC().Add(-1 * time.Hour).Format(time.RFC3339)
	videoItems, err := y.FetchLatestVideo(lastFetchedTime)

	if err != nil {
		return err
	}

	if len(videoItems) > 0 {
		videoMetaDatas, err := y.parse(videoItems)

		if err != nil {
			return err
		}

		err = y.DBClinet.InsertBulk(ctx, videoMetaDatas)
		if err != nil {
			return err
		}

	}
	return nil
}

func (y *YouTubeVideoSyncer) parse(item []manager.VideoItem) ([]*db.VideoMetaData, error) {
	var videoItems []*db.VideoMetaData
	for _, item := range item {
		videoItems = append(videoItems, &db.VideoMetaData{
			Id:               uuid.New().String(),
			VideoId:          item.ID.VideoID,
			Title:            item.Snippet.Title,
			Description:      item.Snippet.Description,
			DefaultThumbnail: item.Snippet.Thumbnails.Default.URL,
			PublishedAt:      y.timeConverter(item.Snippet.PublishedAt),
		})
	}
	return videoItems, nil
}

func (y *YouTubeVideoSyncer) timeConverter(timeInRFC string) int64 {
	parsedTime, err := time.Parse(time.RFC3339, timeInRFC)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return parsedTime.UnixMilli()
}
