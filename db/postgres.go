package db

import (
	"context"
	"gorm.io/gorm"
	"math"
	"time"
)

// TODO: handle error
type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(db *gorm.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (p *PostgresStore) Insert(ctx context.Context, videoMetaData *VideoMetaData) error {
	result := p.db.WithContext(ctx).Create(&videoMetaData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresStore) InsertBulk(ctx context.Context, videoMetaDatas []*VideoMetaData) error {

	result := p.db.WithContext(ctx).CreateInBatches(&videoMetaDatas, len(videoMetaDatas))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresStore) Delete(ctx context.Context, videoId string) error {
	result := p.db.WithContext(ctx).Where("video_id = ?", videoId).Delete(&VideoMetaData{}, videoId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresStore) DeleteBulk(ctx context.Context, videoIds []string) error {
	result := p.db.WithContext(ctx).Where("video_id in (?)", videoIds).Delete(&VideoMetaData{}, videoIds)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresStore) Get(ctx context.Context, videoId string) (*VideoMetaData, error) {
	return nil, nil
}

func (p *PostgresStore) GetPaginated(ctx context.Context, page *Page) ([]*VideoMetaData, error) {
	var videos []*VideoMetaData

	query := p.db.Order("published_at DESC").Limit(page.Size)

	cursorTimeInMillisecond := (int64)(math.MaxInt64)
	if page.LastFetchedTime != "" {
		cursorTime, err := time.Parse(time.RFC3339, page.LastFetchedTime)
		if err == nil {
			cursorTimeInMillisecond = cursorTime.UnixMilli()
		}
	}

	query = query.Where("published_at < ?", cursorTimeInMillisecond)
	result := query.Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}

	return videos, nil
}

func (p *PostgresStore) Search(ctx context.Context, searchQuery string, page *Page) ([]*VideoMetaData, error) {
	var videos []*VideoMetaData

	dbQuery := p.db.Order("published_at DESC").Limit(page.Size)

	cursorTimeInMillisecond := (int64)(math.MaxInt64)
	if page.LastFetchedTime != "" {
		cursorTime, err := time.Parse(time.RFC3339, page.LastFetchedTime)
		if err == nil {
			cursorTimeInMillisecond = cursorTime.UnixMilli()
		}
	}

	if searchQuery != "" {
		likePattern := "%" + searchQuery + "%"
		dbQuery = dbQuery.Where("title ILIKE ? OR description ILIKE ?", likePattern, likePattern)
	}

	dbQuery = dbQuery.Where("published_at < ?", cursorTimeInMillisecond)
	result := dbQuery.Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}

	return videos, nil
}
