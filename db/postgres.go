package db

import (
	"context"
	"gorm.io/gorm"
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
	return nil, nil
}

func (p *PostgresStore) Search(ctx context.Context, query string, page *Page) ([]*VideoMetaData, error) {
	return nil, nil
}
