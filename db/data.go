package db

type VideoMetaData struct {
	id               string `gorm:"primary_key"`
	videoId          string
	title            string
	description      string
	publishedAt      int64
	defaultThumbnail string
}

type Page struct {
	limit  int32
	offset int32
}
